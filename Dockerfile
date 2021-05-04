FROM golang:alpine as build 

WORKDIR /app

COPY . .

RUN go build webserver/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/main /app/main

EXPOSE 8080

ENTRYPOINT ["./main"]