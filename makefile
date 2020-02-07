dev:
	go run main.go

build:
	go build -i -o bin/newsfeeder main.go

run:
	./bin/newsfeeder 