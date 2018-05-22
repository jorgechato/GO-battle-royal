# build stage
FROM golang:1.10 as build-env
WORKDIR /go/src/github.com/jorgechato/battle-royale/
RUN go get -d -v github.com/gorilla/mux
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o battle-royale .

# final stage
FROM scratch
COPY --from=build-env /go/src/github.com/jorgechato/battle-royale/battle-royale /battle-royale
CMD ["/battle-royale"]
