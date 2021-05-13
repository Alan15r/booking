package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/tuloev_alan/booking.core/config"
	"gitlab.com/tuloev_alan/booking.core/handler"
	"gitlab.com/tuloev_alan/booking.core/repository"
	"gitlab.com/tuloev_alan/booking.core/server"
)

const (
	defaultConfigPath = "deployment/config/booking.toml"
)

func main() {
	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	db, err := Connect(&pg.Options{
		Addr:     cfg.Database.Host + ":" + strconv.Itoa(cfg.Database.Port),
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Database: cfg.Database.Db,
	})
	if err != nil {
		log.Fatalf("failed to create a db instance: %v", err)
	}
	defer db.Close()

	hdlr := handler.Handler{
		DB:     &repository.Pg{Db: db},
		Config: cfg,
	}

	router := echo.New()

	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}

	rest.Route()
	rest.Router.Logger.Fatal(router.Start(":3000"))
}

func Connect(op *pg.Options) (*pg.DB, error) {
	conn := pg.Connect(op)

	var ping int
	_, err := conn.QueryOne(pg.Scan(&ping), "SELECT 1")
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the db")
	}
	return conn, nil
}
