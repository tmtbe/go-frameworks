apps = 'app'

.PHONY: run
run: wire
	go run ./cmd
.PHONY: wire
wire:
	go run github.com/google/wire/cmd/wire ./cmd
.PHONY: build
build: clean
	mkdir build
	cp -r resources build
	GOOS=linux GOARCH="amd64" go build -o build/app-linux-amd64 ./cmd;
	GOOS=darwin GOARCH="amd64" go build -o build/app-darwin-amd64 ./cmd;
.PHONY: clean
clean:
	rm -rf build