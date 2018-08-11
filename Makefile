all: build

setup-stamp:
	go get -u github.com/golang/dep/cmd/dep
	touch setup-stamp

setup: setup-stamp

deps-stamp: Gopkg.lock
	dep ensure
	touch deps-stamp

deps: setup-stamp deps-stamp

build: deps
	go build -v

fmt:
	go fmt

test: deps
	go test ./dirtodir

clean:
	rm -rf deps-stamp setup-stamp vendor hashupfs

.PHONY: setup deps cleandeps build
