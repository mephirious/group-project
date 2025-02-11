build:
	@GOOS=linux GOARCH=amd64 go build  -o bin/gateway ./cmd/gateway 

build-win:
	@go build -o bin/gateway.exe  ./cmd/gateway 

run: build
	./bin/gateway

run-win: build-win
	./bin/gateway.exe

test:
	go test -v ./...