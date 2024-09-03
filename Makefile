build:
	go build -o ./bin/app

run: build
	./bin/app

dev:
	go run main.go