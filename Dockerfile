FROM golang:1.23.1-alpine3.20

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN go mod download

CMD ["air", "-c", ".air.toml"]
