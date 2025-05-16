package main

import (
	"log"
	"os"
	"os/signal"
	"point-of-sales-app/adapter"
	"point-of-sales-app/config"
	"point-of-sales-app/internal/routes"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	var errch = make(chan error, 1)
	var shutdown = make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	err := godotenv.Load()
	if err != nil {
		errch <- err
	}

	config.InitConfig()

	if err := adapter.ConnectDB(); err != nil {
		errch <- err
	}

	// errch <- fmt.Errorf("server is running on %s:%s", config.ENV.HOST, config.ENV.PORT)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	g := e.Group("/api")

	routes.InitAuthRoutes(g)
	routes.InitUserRoutes(g)
	routes.InitProductRoutes(g)
	routes.InitCashierRoutes(g)

	go func() {
		log.Printf("Server running on port %s", config.ENV.PORT)
		errch <- e.Start(":" + config.ENV.PORT)
	}()

	select {
	case sig := <-shutdown:
		log.Printf("caught signal %s: shutting down.", sig)

		os.Exit(0)
	case err := <-errch:
		log.Fatalf("caught error: %s", err)
	}
}
