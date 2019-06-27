.PHONY = test fmt
.DEFAULT_GOAL := test

test:
	go test -covermode=count ./...

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;
