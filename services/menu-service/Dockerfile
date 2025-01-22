FROM golang:1.23.1

WORKDIR /services/menu-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o menu-service .

EXPOSE 80

CMD ["./menu-service"]
