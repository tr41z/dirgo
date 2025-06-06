OUTPUT_DIR := bin

TOOL_NAME := dirgo

MAIN_PACKAGE := ./cmd/main/main.go

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(TOOL_NAME)-linux-amd64 $(MAIN_PACKAGE)

# Build for macOS
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(TOOL_NAME)-darwin-amd64 $(MAIN_PACKAGE)

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(TOOL_NAME)-windows-amd64.exe $(MAIN_PACKAGE)

# Build all targets
build-all: build-linux build-darwin build-windows

# Build and run
build-and-run:
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_DIR)/$(TOOL_NAME) $(MAIN_PACKAGE)
	@$(OUTPUT_DIR)/$(TOOL_NAME) -u http://localhost:8080 -w /Users/michael/common.txt -t 5 -s 200

run:
	@./$(OUTPUT_DIR)/$(TOOL_NAME)-darwin-amd64 -u http://localhost:8080 -w /Users/michael/common.txt -t 5 -s 200
