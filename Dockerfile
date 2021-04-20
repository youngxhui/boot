FROM golang:1.16-stretch as build

WORKDIR /app/

COPY . ./boot/
ENV GO111MODULE="on"
ENV GOPROXY=https://goproxy.io,direct
RUN cd boot && ls
RUN cd boot && go mod tidy
RUN cd boot && go build main.go
RUN cd boot && ls
RUN cd boot/gateway && go build main.go

FROM alpine as app
WORKDIR app

COPY --from=build /app/boot/main ./

RUN ls

CMD ["./main"]