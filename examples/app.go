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
