FROM golang:1.22-bookworm

ARG CGO_ENABLED=0

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum /usr/src/app/

RUN go mod download

VOLUME ["/usr/src/app"]

CMD ["go", "version"]
