package main

import (
	"platform-go-challenge/internal/app"
)

func main() {
	application := app.App{}
	application.Init()
	application.Start()
}
