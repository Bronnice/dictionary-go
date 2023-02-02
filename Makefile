build:
	go mod tidy
	go build -o bin/main main.go

run:
	go run main.go
