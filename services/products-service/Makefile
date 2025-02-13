build:
	@GOOS=linux GOARCH=amd64 go build  -o bin/products ./cmd

build-win:
	@go build -o bin/products.exe  ./cmd

run: build
	./bin/products

run-win: build-win
	./bin/products.exe

test:
	go test -v ./...