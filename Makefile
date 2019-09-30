GOBIN=$(shell pwd)/bin

gethapi:
	GOBIN=$(GOBIN) go install cmd/gethapi.go

test:
	go test -cover ./...

nice:
	go fmt ./...

.PHONY: gethapi test nice