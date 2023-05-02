install:
	go mod download
	go mod tidy

install-plugin:
	
	go install github.com/golang/mock/mockgen@latest
	go install github.com/securego/gosec/v2/cmd/gosec@v2.11.0



clean:
	rm -rf generated

run:
	go run cmd/app/main.go

mock:
	mockgen -destination=testing/mocks/Cache.go -package=mocks github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/cache CacheService
	mockgen -destination=testing/mocks/RegisterCourseCheck.go -package=mocks github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/service RegisterCourseCheckService
	
	mockgen -destination=testing/mocks/Client.go -package=mocks github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client Client
	mockgen -destination=testing/mocks/Repository.go -package=mocks github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/repository Repository



test:
	go test -v ./pkg/...

coverage:
	go clean -testcache
	go test -coverpkg=./... -coverprofile cover.out ./...
	go tool cover -html=cover.out
