apps = 'app'
.PHONY: install
install:
	go get github.com/vektra/mockery/v2/.../
	go get github.com/fdaines/arch-go
	go get github.com/go-courier/husky/cmd/husky
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	go mod tidy
	husky init
.PHONY: run
run: clean genApi wire fastRun
fastRun:
	go run ./cmd -f .
.PHONY: wire
wire:
	go run github.com/google/wire/cmd/wire ./cmd
.PHONY: genApi
genApi:
	go run tools/gin-swagger/main.go -f swagger.yaml -A test
.PHONY: build
build: clean
	cp -r resources dist
	GOOS=linux GOARCH="amd64" go build -o dist/app-linux-amd64 ./cmd;
	GOOS=darwin GOARCH="amd64" go build -o dist/app-darwin-amd64 ./cmd;
.PHONY: clean
clean:
	rm -rf dist
	rm -rf gen
	mkdir dist
	mkdir gen
.PHONY: mock
mock:
	mockery --all --output ./gen/mocks
.PHONY: check
check:
	# golangci-lint run
	arch-go
.PHONY: onlyTest
onlyTest:
	go run github.com/google/wire/cmd/wire ./tests
	go test -v  ./internal/app/... -f `pwd` -covermode=count -coverprofile=dist/cover.out
	go test -v  ./tests/... -f `pwd` -covermode=count -coverprofile=dist/cover.out  -json
.PHONY: test
test: clean check genApi mock onlyTest