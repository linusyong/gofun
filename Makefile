all:		depinit dep lint build

depinit:	Gopkg.toml

Gopkg.toml:
	docker run --rm -v ${PWD}:/go/src/$${PWD##*/} -w /go/src/$${PWD##*/} -u $(shell id -u):$(shell id -g) golangdev:alpine dep init

ginkgobootstrap:
	docker run --rm -v ${PWD}:/go/src/$${PWD##*/} -w /go/src/$${PWD##*/} -u $(shell id -u):$(shell id -g) golangdev:alpine ginkgo bootstrap

test:
	docker run --rm -v ${PWD}:/go/src/$${PWD##*/} -w /go/src/$${PWD##*/} -u $(shell id -u):$(shell id -g) golangdev:alpine ginkgo -cover

dep:
	docker run --rm -v ${PWD}:/go/src/$${PWD##*/} -w /go/src/$${PWD##*/} -u $(shell id -u):$(shell id -g) golangdev:alpine dep ensure

build:		ginexample

ginexample:	ginexample.go
	docker run --rm -v ${PWD}:/go/src/$${PWD##*/} -w /go/src/$${PWD##*/} -u $(shell id -u):$(shell id -g) golangdev:alpine go build

dockerimage:
	docker build -t golangdev:alpine .

lint:
	docker run --rm -v ${PWD}:/go/src/$${PWD##*/} -w /go/src/$${PWD##*/} -u $(shell id -u):$(shell id -g) golangdev:alpine golint .

clean:
	rm -rf $${PWD##*/}
	
allclean:
	rm -rf $${PWD##*/} vendor
