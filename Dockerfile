# Use the official Golang image as a parent image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to fetch the dependencies
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Specify the port number which needs to be exposed
EXPOSE 8080

# Command to run the application
CMD ["./main"]
