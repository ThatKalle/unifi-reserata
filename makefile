BINARY_NAME=unifireserata
CURRENT_DIR_WIN=$(shell cd)

buildwin: cleanwin
	go generate
	set GOOS=linux
	set GOARCH=amd64
	go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}-linux
	set GOOS=windows
	set GOARCH=amd64
	go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}-windows.exe

cleanwin:
	go clean
	del /F /Q "resource.syso"
	pushd "${CURRENT_DIR_WIN}\bin" 2>nul && del /F /Q ${BINARY_NAME}-windows.exe
	pushd "${CURRENT_DIR_WIN}\bin" 2>nul && del /F /Q ${BINARY_NAME}-linux

build: clean
	go generate
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}-linux
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}-windows.exe

clean:
	go clean
	rm ./resource.syso || true
	rm ./bin/${BINARY_NAME}-windows.exe || true
	rm ./bin/${BINARY_NAME}-linux || true

dep:
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest

run:
	go run main.go

vet:
	go vet
