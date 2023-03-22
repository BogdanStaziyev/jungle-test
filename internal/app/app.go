package app

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// Echo
	"github.com/labstack/echo/v4"

	// Config
	"github.com/BogdanStaziyev/jungle-test/config"

	// Internal
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/middlewares"
	"github.com/BogdanStaziyev/jungle-test/internal/controller/http/v1"
	"github.com/BogdanStaziyev/jungle-test/internal/database"
	"github.com/BogdanStaziyev/jungle-test/internal/service"

	// External
	"github.com/BogdanStaziyev/jungle-test/pkg/httpserver"
	"github.com/BogdanStaziyev/jungle-test/pkg/jwt"
	"github.com/BogdanStaziyev/jungle-test/pkg/logger"
	"github.com/BogdanStaziyev/jungle-test/pkg/passwords"
	"github.com/BogdanStaziyev/jungle-test/pkg/postgres"
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

	// Create token generator
	tokGen := jwt.NewTokenConstructor(conf.AccessSecret)

	// Create password generator
	passGen := passwords.NewGeneratePasswordHash(conf.Cost)

	// Databases struct of db
	databases := service.Databases{
		AuthRepo:    database.NewAuthRepo(pg),
		ImageRepo:   database.NewImageRepo(pg),
		FileStorage: database.NewStorage(conf.FileStorageLocation),
	}

	// Initialize storage location
	_, err = os.Stat(conf.FileStorageLocation)
	if errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(conf.FileStorageLocation, os.ModePerm)
		if err != nil {
			l.Fatal("storage folder can not be created", "err", err)
		}
	} else if err != nil {
		l.Fatal("storage folder is not available", "err", err)
	}

	// Services struct of all services
	services := v1.Services{
		AuthService:  service.NewAuthService(tokGen, passGen, databases.AuthRepo),
		ImageService: service.NewImageService(databases.ImageRepo, database.NewStorage(conf.FileStorageLocation)),
	}

	// Middleware struct of all middlewares
	middleware := v1.Middleware{
		AuthMiddleware: middlewares.NewMiddleware(conf.AccessSecret),
	}

	// HTTP server start
	handler := echo.New()
	v1.Router(handler, middleware, services, tokGen, l)
	server := httpserver.New(handler, httpserver.Port(conf.ServerPort))

	// Waiting signals
	interrupt := make(chan os.Signal, 1)
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
