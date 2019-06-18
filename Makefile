.PHONY = clean fmt help

main:
	go build main.go

test:
	go test -covermode=count ./...

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;

clean:
	rm -f main