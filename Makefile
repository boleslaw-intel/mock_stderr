test-mock:
	go run main.go > stdout.txt 2> stderr.txt
build:
	go build -o mock-stderr main.go
