apps = 'app'

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
.PHONY: test
test:
	go test -v  ./internal/app/... -f `pwd` -covermode=count -coverprofile=dist/cover.out