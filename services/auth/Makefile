build:
	@GOOS=linux GOARCH=amd64 go build  -o bin/auth ./cmd/auth 

build-win:
	@go build -o bin/auth.exe  ./cmd/auth 

run: build
	./bin/auth

run-win: build-win
	./bin/auth.exe

test:
	go test -v ./...