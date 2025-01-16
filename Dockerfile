# Use the official Go 1.23.2 image
FROM golang:1.23.2-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage to create the final image
FROM alpine:latest  

# Install necessary dependencies for running the Go app
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary and data files from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/data /app/data

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]