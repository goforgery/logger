package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/logger"
)

func main() {
	app := f.CreateApp()
	app.Use(logger.Create())
	app.Listen(3000)
}
