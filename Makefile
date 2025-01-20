.PHONY: build clean

EXECUTABLE_NAME = keil

build:
	go build -o $(EXECUTABLE_NAME)

clean:
	rm -f $(EXECUTABLE_NAME)
