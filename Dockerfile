# Start with a base GoLang image
FROM golang:1.17-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o goscan ./cmd/main

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/goscan /

# Copy the wordlists directory
COPY wordlists /wordlists

# Command to run the executable
ENTRYPOINT ["/goscan"]
