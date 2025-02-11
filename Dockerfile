FROM golang:1.23.5 AS builder

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o server ./cmd/server/main.go

EXPOSE 8080
EXPOSE 8090

CMD ["./server", "-s", "postgres"]