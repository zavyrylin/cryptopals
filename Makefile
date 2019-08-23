.PHONY = test fmt
.DEFAULT_GOAL := test

test:
	go test -v -count=1 -covermode=count ./...

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;
