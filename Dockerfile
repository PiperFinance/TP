ARG GO_VERSION=1.21

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
ENV PORT=8000
COPY  ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download

COPY ./src .
RUN go build -o ./app TP

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api \
    && rm -rf /var/bs/log/ | true \ 
    && mkdir -p /var/bs/log/ \ 
    && touch /var/bs/log/err.log \ 
    && touch /var/bs/log/debug.log \
    && rm -rf /var/cache/apk/*


WORKDIR /api
COPY --from=builder /api/app .
EXPOSE 8080

ENTRYPOINT ["./app"]

