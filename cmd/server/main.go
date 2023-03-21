package main

import (
	"github.com/BogdanStaziyev/jungle-test/config"
	"github.com/BogdanStaziyev/jungle-test/internal/app"
)

func main() {
	// Initialize configuration
	conf := config.GetConfiguration()

	// Run application
	app.Run(conf)
}
