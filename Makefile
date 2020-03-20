default:
	go build ./...

test:
	go test ./...

.PHONY: default test
