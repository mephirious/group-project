build:
	@GOOS=linux GOARCH=amd64 go build  -o bin/blogs ./cmd

build-win:
	@go build -o bin/blogs.exe  ./cmd

run: build
	./bin/blogs

run-win: build-win
	./bin/blogs.exe

test:
	go test -v ./...