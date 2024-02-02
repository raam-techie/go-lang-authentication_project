# Use a Golang base image
FROM golang:1.21.6 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a minimal base image to reduce image size
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

COPY --from=builder /app/app .

# Command to run the executable
CMD ["./app"]
