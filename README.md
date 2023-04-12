# pprof

`pprof` is a Go package that provides a set of HTTP handlers for profiling Go programs. It is built on top of the standard `net/http/pprof` package and is designed to work with the `lightning` web framework.

## Installation

```bash
go get github.com/lightning-contrib/pprof
```

## Usage

Here is an example of how to use this package:

```go
package main

import (
	"github.com/go-labx/lightning"
	"github.com/lightning-contrib/pprof"
)

func main() {
	app := lightning.DefaultApp()
	pprof.Register(app)
	app.Run()
}
```

By default, the package's HTTP handlers are registered under the `/debug/pprof` prefix. You can customize the prefix by passing an optional argument to the `Register` function

```go
pprof.Register(app, "/admin/debug/pprof")
```

Once the package's HTTP handlers are registered, you can access them using a web browser or a command-line tool like `go tool pprof`. Here are some examples:

- To view the heap profile of your Go program, visit http://127.0.0.1:6789/debug/pprof/heap?debug=1 in your web browser.
- To view the goroutine profile of your Go program, visit http://127.0.0.1:6789/debug/pprof/goroutine?debug=1 in your web browser.
- To view the block profile of your Go program, visit http://127.0.0.1:6789/debug/pprof/block?debug=1 in your web browser.
- To view the mutex profile of your Go program, visit http://127.0.0.1:6789/debug/pprof/mutex?debug=1 in your web browser.
- To view the symbol table of your Go program, visit http://localhost:6789/debug/pprof/symbol?debug=1 in your web browser.

Note that the above URLs assume that your lightning application is running on `localhost` and listening on port `6789`. If your application is running on a different host or port, you need to adjust the URLs accordingly.

## API Documentation

For detailed API documentation and usage examples, please refer to the [documentation](https://pkg.go.dev/github.com/lightning-contrib/pprof).

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](https://github.com/lightning-contrib/pprof/blob/main/CONTRIBUTING.md) for more information.

## License

This package is licensed under the [MIT License](https://github.com/lightning-contrib/pprof/blob/main/LICENSE). 
