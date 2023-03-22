package main

import (
	// Config
	"github.com/BogdanStaziyev/jungle-test/config"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/app"
)

func main() {
	// Initialize configuration
	conf := config.GetConfiguration()

	// Run application
	app.Run(conf)
}
