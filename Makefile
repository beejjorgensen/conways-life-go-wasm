hello.wasm: hello.go
	GOOS=js GOARCH=wasm go build -o hello.wasm