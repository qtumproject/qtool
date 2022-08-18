FROM golang:1.14-alpine

RUN apk add --no-cache make gcc musl-dev git

WORKDIR /qtool-api
COPY ./ /qtool

RUN go build \
        -o /qtool//bin /qtool-api/...

EXPOSE 8080

ENTRYPOINT [ "qtool-api" ]