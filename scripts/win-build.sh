mkdir -p bin
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o bin/lemon_cli.exe github.com/entropy-platform/lemon/cli
