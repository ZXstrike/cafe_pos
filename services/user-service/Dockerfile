FROM golang:1.23.1

WORKDIR /services/user-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service .

EXPOSE 80

CMD ["./user-service"]
