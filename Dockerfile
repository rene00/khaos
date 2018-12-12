FROM golang:alpine as builder

RUN apk update && apk add --no-cache git ca-certificates libc6-compat gcc g++

RUN adduser -D -g '' khaos

COPY . $GOPATH/src/khaos

WORKDIR $GOPATH/src/khaos

RUN GO111MODULE=on go mod download

RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -s -extldflags "-static"' -o /go/bin/khaos ./cmd/khaos

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /go/bin/khaos /go/bin/khaos

USER khaos

ENTRYPOINT ["/go/bin/khaos", "start", "--database-uri=/tmp/khaos.db"]
