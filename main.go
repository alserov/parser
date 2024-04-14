package main

import (
	_ "github.com/alserov/parser/docs"
	"github.com/alserov/parser/internal/app"
	"github.com/alserov/parser/internal/config"
)

// @title Parser API
// @version 1.0

func main() {
	app.MustStart(config.MustLoad())
}
