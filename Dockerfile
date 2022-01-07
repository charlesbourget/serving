# syntax=docker/dockerfile:1
FROM golang:latest as build

WORKDIR /go/src/github.com/cbourget/serving/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/github.com/cbourget/serving/app ./
CMD ["./app", "-d", "/mount", "-f", "json", "-p", "8080"]
