FROM golang:1.21.1-alpine as base
WORKDIR /app/go
COPY . .
RUN go mod download
RUN go mod verify

# Copy the local code and .env file to the container
COPY .env .env

# Create the public/uploads folder
RUN mkdir -p public/uploads

# Install Swag
# RUN go get -u github.com/swaggo/swag/cmd/swag

# Run Swag to generate Swagger documentation
# RUN swag init --pd -g ./cmd/main/main.go -o ./docs

# Build the Go application
RUN go build -o cmd/main/main ./cmd/*/*.go

# Expose the port the application will run on
EXPOSE 9545

ENV TZ="Asia/Bangkok"

# Command to run the executable
CMD ["./cmd/main/main"]