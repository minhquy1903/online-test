package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	authHttp "github.com/minhquy1903/online-test/auth-service/internal/auth/delivery/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/online-test/auth-service/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	db     *gorm.DB
	logger *logrus.Logger
	ready  chan bool
}

func NewServer(cfg *config.Config, db *gorm.DB, logger *logrus.Logger, ready chan bool) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, logger: logger, ready: ready}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         ":" + s.cfg.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalln("Error starting Server: ", err)
		}
	}()

	v1 := s.echo.Group("/api/v1")

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup)

	if s.ready != nil {
		s.ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.logger.Fatalln("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}