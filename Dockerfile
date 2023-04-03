<<<<<<< HEAD
FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["./main","run"]
=======
# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download Go modules
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Run the binary program produced by `go build`
CMD ["./main"]
>>>>>>> ef140560fa3717d83e953e94a39105a691a2f442
