ARG GO_VERSION=1.18
ARG ALPINE_VERSION=3.16

FROM golang:${GO_VERSION}-alpine as builder

RUN apk add --no-cache make gcc musl-dev git

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./qtool-api ./qtool-api
COPY ./pkg  ./pkg
RUN cd qtool-api && go build -o /app/bin/qtool-api

FROM alpine:${ALPINE_VERSION}

WORKDIR /app/bin
COPY --from=builder /app/bin/qtool-api .
EXPOSE 8080
CMD [ "./qtool-api" ]