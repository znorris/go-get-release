FROM golang:1-alpine

RUN apk add --no-cache git ca-certificates

WORKDIR /go/src/go-get-release
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

VOLUME ["/tmp/downloads"]
WORKDIR /tmp/downloads

ENTRYPOINT ["/go/bin/go-get-release"]
