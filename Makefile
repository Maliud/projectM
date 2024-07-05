build:
	go build -o ./bin/projectm

run: build
	./bin/projectm

test:
	go test ./...