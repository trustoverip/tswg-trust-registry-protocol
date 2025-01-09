# Use the official Golang image as the base image
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application code
COPY . .

# Build the application binary
RUN go build -o server .

# Use a minimal image for the final container
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Copy any necessary static files like the Swagger and Redoc files
COPY ./api/trqp.yaml /app/api/trqp.yaml
COPY ./redoc.html /app/redoc.html
COPY ./data/registry.json /app/data/registry.json

# Expose the port your application runs on
EXPOSE 8082

# Command to run the application
CMD ["./server"]
