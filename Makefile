.PHONY : test build clean format

build:
	go build github.com/telkomdev/baper/cmd/baper

test:
	go test ./...

test-verbose:
	go test -v ./...

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

clean:
	rm baper *.txt