FROM golang:alphine

LABEL maintainer="Siang <dysiang0429@gmail.com>"

RUN apk update && apk add --no-cache git

RUN go get github.com/lib/pq