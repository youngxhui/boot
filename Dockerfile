FROM golang:1.16-stretch as build

WORKDIR /tmp/

COPY . ./boot/
ENV GO111MODULE="on"
ENV GOPROXY=https://goproxy.io,direct
RUN cd boot && go mod tidy
RUN cd boot && go build main.go
RUN cd boot/gateway && go build main.go

FROM alpine as app

WORKDIR /app
COPY from=0 /tmp/boot/main ./

CMD ./main

FROM alpine as gateway

WORKDIR /app
COPY from=build /tmp/boot/gateway/main ./

CMD ./main
