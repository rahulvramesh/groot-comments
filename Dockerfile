#builder
FROM golang:latest AS builder
ADD . $GOPATH/src/github.com/rahulvramesh/groot-comments
WORKDIR $GOPATH/src/github.com/rahulvramesh/groot-comments
RUN go get -u github.com/golang/dep/cmd/dep
RUN make server

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /server ./
RUN chmod +x ./server
ENTRYPOINT ["./server"]
EXPOSE 8005