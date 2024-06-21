# Use the official Golang image as a build stage
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /blogpost

# Use a smaller base image for the final build


# Set the working directory inside the container

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["/blogpost"]
