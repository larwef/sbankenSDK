all: fmt test

fmt:
	go fmt ./...
	golint ./...

test:
	go test ./...

doc:
	godoc -http=":6060"
