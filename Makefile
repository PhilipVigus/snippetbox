.PHONY: build run test clean

build:
	docker compose up --build --no-deps app

test:
	go test ./...

clean:
	rm -f app
