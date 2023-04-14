install:
	go mod download
	go mod tidy

install-plugin:
	
	go install github.com/golang/mock/mockgen@latest


clean:
	rm -rf generated

run:
	go run cmd/app/main.go

gen:
	mockgen -destination=testing/mocks/Repository.go -package=mocks github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/repository Repository


test:
	go test -v ./pkg/...

coverage:
	go clean -testcache
	go test -coverpkg=./... -coverprofile cover.out ./...
	go tool cover -html=cover.out
