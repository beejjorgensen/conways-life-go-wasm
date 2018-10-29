hello.wasm: hello.go life/life.go
	GOOS=js GOARCH=wasm go build -o hello.wasm