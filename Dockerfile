FROM golang:1.15.0-alpine3.12 as builder
RUN apk update && apk upgrade && \
    apk --no-cache --update add git make && \
    go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/wincentrtz/gobase
COPY domains/user/mocks .
RUN dep ensure -v && go build -o engine app/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    mkdir /app && mkdir gobase
WORKDIR /gobase
EXPOSE 8080
COPY --from=builder /go/src/github.com/wincentrtz/gobase /app
CMD /app/engine