package main

import (
	"github.com/kbondar17/go-server/internal/app"
)

func main() {
	app := app.NewApp()
	app.Run()

}
