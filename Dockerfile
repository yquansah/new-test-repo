FROM golang:1.21-alpine

WORKDIR /app

COPY go.* .
COPY main.go .

RUN go build -o /sample

CMD ["/sample"]
