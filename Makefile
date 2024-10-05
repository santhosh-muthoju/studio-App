APP_NAME = studio-app
BINARY = bin/$(APP_NAME)

build:
	go build -o $(BINARY) ./app

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)
