# Use the official Golang image as the base image with Go 1.21.1
FROM golang:1.21.1

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code and go.mod/go.sum files into the container
COPY go.mod .
COPY go.sum .

# Download and install the Go module dependencies
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Go application
RUN go build -o scorecard-go

# Expose the port that your application listens on (replace with your application's port)
EXPOSE 8080

# Set the command to run your Go application
CMD ["./scorecard-go"]
