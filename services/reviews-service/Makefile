build:
	@GOOS=linux GOARCH=amd64 go build  -o bin/reviews ./cmd

build-win:
	@go build -o bin/reviews.exe  ./cmd

run: build
	./bin/reviews

run-win: build-win
	./bin/reviews.exe

test:
	go test -v ./...