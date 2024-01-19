build:
	go mod tidy
	go build -ldflags "-s -w" -o dist/ .

run:
	go mod tidy
	go install github.com/cosmtrek/air@latest
	air

compile:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/echodemo-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/echodemo-macos-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/echodemo-windows-amd64.exe .