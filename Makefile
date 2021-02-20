build:
	go build -ldflags "-s -w" -o news main.go

run:
	go run main.go