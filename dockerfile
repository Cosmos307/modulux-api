FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy the Go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o modulux .

# Set execute permissions for the binary
RUN chmod +x ./modulux

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./modulux"]