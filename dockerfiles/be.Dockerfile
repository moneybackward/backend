FROM golang:1.21.3-alpine3.18

WORKDIR /go/src/app

COPY ./.. .

RUN go mod download

EXPOSE 8000

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

RUN go build -o moneybackward-be
