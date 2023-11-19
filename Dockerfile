FROM golang:latest

ENV GO111MODULE=on

COPY . /app/

WORKDIR /app

RUN mkdir -p /app/logs

RUN go mod download

ENTRYPOINT ["/bin/bash", "-c", "go run main.go"]