# Conway's Life in Go and Web Assembly

Experiments with Go and WASM. I'm fairly beginner at both, so feedback
is welcome.

## Install and run

A Makefile is included to help with the build.

```
make
```

And then run a webserver that's aware of the `application/wasm` MIME
type. If you don't have one handy, here's a silent one-liner for use
with [goexec](https://github.com/shurcooL/goexec) listening on
`http://localhost:8000/`:

```
goexec 'http.ListenAndServe(":8000", http.FileServer(http.Dir(".")))'
```

## Demo

TBD

## References

* [Go/WASM wiki article](https://github.com/golang/go/wiki/WebAssembly)
* [syscall/js docs](https://golang.org/pkg/syscall/js/)