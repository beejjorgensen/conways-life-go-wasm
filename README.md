# Conway's Life in Go and Web Assembly

Experiments with Go and WASM. I'm fairly beginner at both, so feedback
is welcome.

## Install and run

Go v1.11 or higher is required.

A Makefile is included to help with the build.

```
make
```

This'll put a `conlife.wasm` file in the `docs/` directory.

And then run a webserver in that directory. The server has to be aware
of the `application/wasm` MIME type. If you don't have one handy, here's
a silent one-liner for use with
[goexec](https://github.com/shurcooL/goexec) listening on
`http://localhost:8000/`:

```
goexec 'http.ListenAndServe(":8000", http.FileServer(http.Dir(".")))'
```

## Demo

TBD

## References

* [Go/WASM wiki article](https://github.com/golang/go/wiki/WebAssembly)
* [syscall/js docs](https://golang.org/pkg/syscall/js/)