package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/online-test/api-gateway/pkg/auth"
	"github.com/minhquy1903/online-test/api-gateway/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {

    echo := echo.New()
    cfg, err := config.LoadConfig()
    logger := logrus.Logger{}
    var ready chan bool

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    server := &http.Server{
		Addr:         ":"   + cfg.Port,
		WriteTimeout: 15    * time.Second,
		ReadTimeout:  15    * time.Second,
	}

	go func() {
		logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", cfg.Port)

		if err := echo.StartServer(server); err != nil {
			logger.Fatalln("Error starting Server: ", err)
		}

	}()

	authGroup := echo.Group("/api/v1")

	auth.RegisterRoutes(authGroup, &cfg)

	if ready != nil {
		ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Fatalln("Server Exited Properly")
	echo.Server.Shutdown(ctx)
}
