FROM golang:alpine

RUN apk --no-cache add git && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u golang.org/x/lint/golint && \
    go get github.com/onsi/ginkgo/ginkgo && \
    go get github.com/onsi/gomega/... && \
    mkdir -p /.cache && \
    chmod -R 777 /go /.cache
    
