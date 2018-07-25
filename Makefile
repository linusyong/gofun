all:		depinit dep lint build

depinit:	Gopkg.toml

Gopkg.toml:
	docker run --rm -v ${PWD}:/go/src/myapp -w /go/src/myapp golangdep:alpine dep init

dep:
	docker run --rm -v ${PWD}:/go/src/myapp -w /go/src/myapp golangdep:alpine dep ensure -update

build:		app

app:		main.go
	docker run --rm -v ${PWD}:/go/src/myapp -w /go/src/myapp golangdep:alpine go build -o app

dockerimage:
	docker build -t golangdep:alpine .

lint:
	docker run --rm -v ${PWD}:/go/src/myapp -w /go/src/myapp golangdep:alpine golint .

clean:
	rm -rf app vendor Gopkg*
