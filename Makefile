
clean:
	rm -rf push-relay

test: clean
	go test ./... -count=1 -race -cover

build: test
	go build ./...

