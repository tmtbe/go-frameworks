apps = 'app'
.PHONY: install
install:
	go get github.com/vektra/mockery/v2/.../
	go get github.com/fdaines/arch-go
	go get github.com/go-courier/husky/cmd/husky
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	go mod tidy
	go get github.com/google/wire/cmd/wire
	husky init
.PHONY: run
run: wire
	go run ./cmd -f .
.PHONY: wire
wire:
	go run github.com/google/wire/cmd/wire ./cmd
.PHONY: build
build: clean
	mkdir dist
	cp -r resources dist
	GOOS=linux GOARCH="amd64" go build -o dist/app-linux-amd64 ./cmd;
	GOOS=darwin GOARCH="amd64" go build -o dist/app-darwin-amd64 ./cmd;
.PHONY: clean
clean:
	rm -rf dist
	mkdir dist
.PHONY: mock
mock: clean
	mockery --all --output ./tests/mocks
.PHONY: check
check:
	golangci-lint run
	arch-go
.PHONY: onlyTest
onlyTest:
	go run github.com/google/wire/cmd/wire ./tests
	go test -v  ./internal/app/... -f `pwd` -covermode=count -coverprofile=dist/cover.out
	go test -v  ./tests/... -f `pwd` -covermode=count -coverprofile=dist/cover.out
.PHONY: test
test: check mock onlyTest