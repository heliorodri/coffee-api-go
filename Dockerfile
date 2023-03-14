# Use the official Golang image as the base
FROM golang:1.16-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o app .

# Use a lightweight image for the production container
FROM alpine:3.13

# Set the working directory
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=build /app/app .

# Set the command to run when the container starts
CMD ["./app"]
