export GO111MODULE=on
BINARY_NAME=bin/pokecli
MAIN=bin/main.go

all: deps build
install:
	go install $(MAIN)
build:
	go build -o $(BINARY_NAME) $(MAIN) && chmod +x $(BINARY_NAME)
test:
	go test -v ./...
clean:
	go clean
	rm -f $(BINARY_NAME)
deps:
	go build -v ./...
upgrade:
	go get -u
run:
	$(BINARY_NAME) $@
