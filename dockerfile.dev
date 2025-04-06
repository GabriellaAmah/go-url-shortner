FROM golang:1.23 AS build-stage

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 7090

CMD  ["air"]