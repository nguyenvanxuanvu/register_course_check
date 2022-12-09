install:
	go mod download
	go mod tidy


clean:
	rm -rf generated

run:
	go run cmd/app/main.go

