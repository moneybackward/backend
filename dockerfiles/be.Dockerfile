FROM golang:1.21.3-alpine3.18

WORKDIR /go/src/app

COPY ./.. .

RUN echo $(ls -la)
RUN go mod download

EXPOSE 8000

RUN swag init

RUN go build -o moneybackward-be
