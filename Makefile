.PHONY: build test clean dist dev-run doc-server

EXECUTABLE_NAME = keil

build:
	go build -o $(EXECUTABLE_NAME)

test:
	go test ./...

clean:
	rm -f $(EXECUTABLE_NAME)
	rm -f out.json
	rm -f -r dist

dist:
	GOOS=darwin GOARCH=arm64 go build -o ./dist/$(EXECUTABLE_NAME)-darwin-arm64
	GOOS=linux GOARCH=amd64 go build -o ./dist/$(EXECUTABLE_NAME)-linux-amd64
	GOOS=windows GOARCH=amd64 go build -o ./dist/$(EXECUTABLE_NAME)-windows-amd64

dev-run:
	go fmt
	make build
	./$(EXECUTABLE_NAME) merge -d testdata --enableIndentation
	rm -f $(EXECUTABLE_NAME)

doc-server:
	godoc -http=:6060