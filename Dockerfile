FROM golang:alpine

RUN apk --no-cache add git; \
    wget -qO - https://raw.githubusercontent.com/golang/dep/master/install.sh | sh; \
    go get -u golang.org/x/lint/golint
