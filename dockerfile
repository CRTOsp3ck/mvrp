# Use the official Golang image as a base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o mvrp-api ./cmd/app/

# Set environment variables
ENV PORT=6900
ENV DB_HOST=db
ENV DB_PORT=5432

# Expose the application's port
EXPOSE ${PORT}

# Command to run the application
CMD ["./mvrp-api"]
