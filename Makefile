.PHONY: all build test clean docker install uninstall

all: clean test build docker install

build:
	go build -o ./build/sixSense main.go

test:
	go test -v ./...

clean:
	rm -rf ${GOPATH}/pkg/github.com/ttante/sixSense

docker:
	docker build -t sixsense .

install:
	cp -r ./build/ /usr/local/bin/

uninstall:
	rm /usr/local/bin/sixSense
