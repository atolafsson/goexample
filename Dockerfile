FROM golang:1.19.3

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy