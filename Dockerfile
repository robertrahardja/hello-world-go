# Start from the official Go image
FROM golang:1.21-alpine AS builder

# Install git and templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Generate templ files
RUN templ generate

# Build the application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy any other necessary files (like templates, static files, etc.)
# COPY --from=builder /app/templates ./templates
# COPY --from=builder /app/static ./static

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]
