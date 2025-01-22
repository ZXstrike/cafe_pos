FROM golang:1.23.1

WORKDIR /api-gateway

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-gateway .

EXPOSE 80

CMD ["./api-gateway"]

# The Dockerfile for the API Gateway service is similar to the Dockerfile for the user-service application.

# The Dockerfile starts by using the official Golang image version 1.23.1 as the base image.

# The working directory is set to /go/src/github.com/zxstrike/api-gateway.

# The go.mod and go.sum files are copied to the working directory, and go mod download is run to download the Go module dependencies.

# The entire application code is copied to the working directory.

# The Go application is built, and the binary is output as api-gateway.

# Port 80 is exposed for the application.

# The command to run the application binary ./api-gateway is specified.

# The Dockerfile for the API Gateway service is now complete. You can now build the Docker image for the API Gateway service.

