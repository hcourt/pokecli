export GO111MODULE=on
BINARY=pokecli
MAIN=bin/pokecli.go

all: deps install
install:
	go install $(MAIN)
build:
	go build $(MAIN)
test:
	go test -v ./...
clean:
	go clean ./...
deps:
	go build -v ./...
upgrade:
	go get -u
run:
	$(BINARY_NAME)
