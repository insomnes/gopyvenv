build:
	go build -o "bin/" cmd/gopyvenv/gopyvenv.go

test:
	go test -v ./...

install:
	/usr/bin/env bash ./install.sh
