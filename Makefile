fmt:
	go fmt ./...

test:
	go test -v ./...

imports:
	goimports -w .

.PHONY: fmt test imports

all: test fmt imports

.PHONY: all