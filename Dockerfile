FROM golang:1.18-alpine

RUN apk add --no-cache make gcc musl-dev git

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./qtool-api ./qtool-api
COPY ./pkg  ./pkg
RUN cd qtool-api && go build -o /app/bin/qtool-api

EXPOSE 8080

CMD [ "/app/bin/qtool-api" ]