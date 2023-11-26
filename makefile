BINARY_NAME=unifireserata
CURRENT_DIR_WIN=$(shell cd)

buildwin: cleanwin
	go generate
	set GOOS=linux
	set GOARCH=amd64
	go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}
	set GOOS=windows
	set GOARCH=amd64
	go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}.exe

cleanwin:
	go clean
	del /F /Q "resource.syso"
	pushd "${CURRENT_DIR_WIN}\bin" 2>nul && del /F /Q ${BINARY_NAME}
	pushd "${CURRENT_DIR_WIN}\bin" 2>nul && del /F /Q ${BINARY_NAME}.exe

build: clean
	go generate
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME}.exe

clean:
	go clean
	rm ./resource.syso || true
	rm ./bin/${BINARY_NAME} || true
	rm ./bin/${BINARY_NAME}.exe || true

dep:
	go mod download
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
	go install github.com/goreleaser/goreleaser@latest

run:
	go run main.go

vet:
	go vet
