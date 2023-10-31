package main

import (
	"go.uber.org/fx"
	"person-service/internal"
	"person-service/internal/router"
	"runtime"
)

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	app := fx.New(
		internal.Modules,
		router.EntryPoint,
	)
	app.Run()
}
