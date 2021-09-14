FROM golang:1.16-alpine AS build-env

WORKDIR /go/src/verify-my-test

RUN apk update
RUN apk add git

COPY ./docker/go/entrypoint.sh /root/
COPY ./example.env ./.env
RUN chmod 755 /root/entrypoint.sh

COPY go.mod .
COPY go.sum .
RUN go mod download

# Project files
COPY . .
COPY ./docker/go/dbconf.yml ./db/dbconf.yml
RUN go build -o verify-my-test

RUN chmod +x ./verify-my-test



FROM alpine:3.10
WORKDIR /app
COPY --from=build-env /go/src/verify-my-test/.env /app/.env
COPY --from=build-env /go/src/verify-my-test/doc/swagger.json /app/doc/swagger.json
COPY --from=build-env /go/src/verify-my-test/verify-my-test /app/
EXPOSE 8989
ENTRYPOINT ./verify-my-test server