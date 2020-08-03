FROM golang:alpine as builder

WORKDIR /src

ADD pullme.go /src

RUN apk update && apk add --no-cache git && \
    GOOS=linux go build -ldflags="-w -s" pullme.go && \
    adduser -D -g '' puller

FROM alpine:3.8

COPY --from=builder /src/pullme /app/pullme
COPY --from=builder /etc/passwd /etc/passwd

USER puller

WORKDIR /app

ENTRYPOINT [ "/app/pullme" ]
