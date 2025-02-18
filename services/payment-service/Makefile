build:
	@GOOS=linux GOARCH=amd64 go build  -o bin/payment ./cmd

build-win:
	@go build -o bin/payment.exe  ./cmd

run: build
	./bin/payment

run-win: build-win
	./bin/payment.exe

test:
	go test -v ./...