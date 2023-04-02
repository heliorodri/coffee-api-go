FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]

