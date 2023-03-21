package app

import (
	"fmt"
	"github.com/BogdanStaziyev/jungle-test/config"
	v1 "github.com/BogdanStaziyev/jungle-test/internal/controller/http/v1"
	"github.com/BogdanStaziyev/jungle-test/pkg/httpserver"
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
	"github.com/BogdanStaziyev/jungle-test/pkg/postgres"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
)

func Run(conf config.Configuration) {
	l := logger.New(conf.LogLevel)

	// Start migrations
	if err := Migrate(conf); err != nil {
		l.Fatal(fmt.Errorf("unable to apply migrations: %s", err))
	}

	pgURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		conf.DatabaseUser, conf.DatabasePassword, conf.DatabaseHost, conf.DatabasePort, conf.DatabaseName,
	)

	// Connect to database
	pg, err := postgres.New(pgURL)
	if err != nil {
		l.Fatal(fmt.Errorf("unable to make postgreSQL connection: %s", err))
	}
	defer pg.Close()

	// Create password generator
	//passGen := passwords.NewGeneratePasswordHash(conf.Cost)

	// Databases struct of db
	//databases := service.Databases{}

	// Services struct of all services
	services := v1.Services{}

	// HTTP server start
	handler := echo.New()
	server := httpserver.New(handler, httpserver.Port(conf.ServerPort))

	// Waiting signals
	interrupt := make(chan os.Signal, 1)
	v1.Router(handler, services, l)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Error("Signal interrupt error: " + s.String())
	case err = <-server.Notify():
		l.Error("Server notify", "err", err)
	}

	// Shutdown server
	err = server.Shutdown()
	if err != nil {
		l.Error("Server shutdown: ", "err", err)
	}
}
