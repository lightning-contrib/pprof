package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-labx/lightning"
)

const defaultPrefix = "/debug/pprof"

// wrapF is a helper function for wrapping http.HandlerFunc and returns a Gin middleware.
func wrapF(f http.HandlerFunc) lightning.HandlerFunc {
	return func(c *lightning.Context) {
		c.SkipFlush()
		c.SetStatus(http.StatusOK)
		f(c.Res, c.Req)
	}
}

// wrapH is a helper function for wrapping http.Handler and returns a Gin middleware.
func wrapH(h http.Handler) lightning.HandlerFunc {
	return func(c *lightning.Context) {
		c.SkipFlush()
		c.SetStatus(http.StatusOK)
		h.ServeHTTP(c.Res, c.Req)
	}
}

func Register(app *lightning.Application, prefixOptions ...string) {
	prefix := defaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}

	group := app.Group(prefix)
	group.Get("/", wrapF(pprof.Index))
	group.Get("/cmdline", wrapF(pprof.Cmdline))
	group.Get("/profile", wrapF(pprof.Profile))
	group.Post("/symbol", wrapF(pprof.Symbol))
	group.Get("/symbol", wrapF(pprof.Symbol))
	group.Get("/trace", wrapF(pprof.Trace))
	group.Get("/allocs", wrapH(pprof.Handler("allocs")))
	group.Get("/block", wrapH(pprof.Handler("block")))
	group.Get("/goroutine", wrapH(pprof.Handler("goroutine")))
	group.Get("/heap", wrapH(pprof.Handler("heap")))
	group.Get("/mutex", wrapH(pprof.Handler("mutex")))
	group.Get("/threadcreate", wrapH(pprof.Handler("threadcreate")))
}
