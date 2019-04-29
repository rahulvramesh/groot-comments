# Builder
FROM golang:latest as builder

RUN apk update && apk upgrade && \
    apk --update add git gcc make && \
    go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/rahulvramesh/groot-comments

COPY . .

RUN make server

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8005

COPY --from=builder /go/src/github.com/rahulvramesh/groot-comments/server /app

CMD /app/server