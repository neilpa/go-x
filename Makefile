all: default windows

default:
	go build ./...

windows:
	GOOS=windows go build ./...

test:
	go test ./...

.PHONY: all default windows test
