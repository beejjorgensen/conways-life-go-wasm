conlife.wasm: conlife.go life/life.go
	GOOS=js GOARCH=wasm go build -o conlife.wasm