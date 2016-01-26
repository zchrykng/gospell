
all: install lint test

install:
	go get ./...
	go install ./...

lint:
	golint ./...
	go vet ./...
	find . -name '*.go' | xargs gofmt -w -s

test:
	go test .
	misspell *.md *.go

clean:
	rm -f *~
	go clean ./...

ci: install lint test

docker-ci:
	docker run --rm \
		-e COVERALLS_REPO_TOKEN=$COVERALLS_REPO_TOKEN \
		-v $(PWD):/go/src/github.com/client9/misspell \
		-w /go/src/github.com/client9/misspell \
		nickg/golang-dev-docker \
		make ci

.PHONY: ci docker-ci