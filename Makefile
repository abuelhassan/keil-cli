.PHONY: build clean dist

EXECUTABLE_NAME = keil

build:
	go build -o $(EXECUTABLE_NAME)

clean:
	rm -f $(EXECUTABLE_NAME)
	rm -f out.json
	rm -f -r dist

dist:
	GOOS=darwin GOARCH=arm64 go build -o ./dist/$(EXECUTABLE_NAME)-darwin-arm64
	GOOS=linux GOARCH=amd64 go build -o ./dist/$(EXECUTABLE_NAME)-linux-amd64
	GOOS=windows GOARCH=amd64 go build -o ./dist/$(EXECUTABLE_NAME)-windows-amd64
