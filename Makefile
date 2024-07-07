build:
	@mkdir -p bin
	@go build -o bin/goscan ./cmd/main

run: build
	@./bin/goscan -u http://localhost:8080 -w ./bin/common.txt -t 5 -s 200 -o results.txt