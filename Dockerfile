FROM golang:1.21.1-alpine as base

# Set the working directory and change ownership to the new user
WORKDIR /app/go

# Copy the application code
COPY . .
RUN go mod download
RUN go mod verify

# Create the public/uploads folder
RUN mkdir -p public/uploads

# Install Swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest
# Run Swag to generate Swagger documentation
RUN swag init --pd -g ./cmd/main/main.go -o ./docs

# Build the Go application
# RUN go build -o cmd/main/main ./cmd/*/*.go
RUN GOOS=linux GOARCH=amd64 go build -o cmd/main/main ./cmd/*/*.go

# Use a lightweight image as the final base image
FROM alpine:3.14

# Install curl
RUN apk --no-cache add curl

# Set the working directory
WORKDIR /app/go

# Copy the binary from the base image
COPY --from=base /app/go/cmd/main/main .

# Expose the port the application will run on
EXPOSE 9545

ENV GIN_MODE=release
ENV TZ=Asia/Bangkok

# Health check endpoint
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD curl --fail http://localhost:9545/health || exit 1

# Command to run the executable
CMD ["./main"]