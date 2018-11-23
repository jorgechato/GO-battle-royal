FROM golang:alpine as builder

COPY . $GOPATH/src/github.com/jorgechato/battle-royal
WORKDIR $GOPATH/src/github.com/jorgechato/battle-royal

RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="${LDFLAGS}" -o /go/bin/server server.go


FROM scratch

COPY --from=builder /go/bin/server /opt/server

ENTRYPOINT ["/opt/server"]
