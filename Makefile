default:
	go build ./...
	GOOS=windows go build ./win

test:
	go test ./...

.PHONY: default test
