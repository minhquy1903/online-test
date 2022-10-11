package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/minhquy1903/online-test/auth-service/config"
	"github.com/minhquy1903/online-test/auth-service/internal/auth/delivery/pb"
	"github.com/minhquy1903/online-test/auth-service/internal/auth/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	db     *gorm.DB
	logger *logrus.Logger
	ready  chan bool
}

func NewServer(cfg *config.Config, db *gorm.DB, logger *logrus.Logger, ready chan bool) *Server {
	return &Server{cfg: cfg, db: db, logger: logger, ready: ready}
}

func (s *Server) Run() error {
	// server := &http.Server{
	// 	Addr:         ":" + s.cfg.Server.Port,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// go func() {
	// 	s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", s.cfg.Server.Port)
	// 	if err := s.echo.StartServer(server); err != nil {
	// 		s.logger.Fatalln("Error starting Server: ", err)
	// 	}
	// }()

	// Init repositories
	
	// Init useCases

	svc := service.NewAuthService(s.db, s.cfg)

	h := rpc{}

	lis, err := net.Listen("tcp", s.cfg.Server.Port)

	if err != nil {
        s.logger.Fatalln("Failed to listing:", err)
    }

    grpcServer := grpc.NewServer()

    pb.RegisterAuthServiceServer(grpcServer, h)

    if err := grpcServer.Serve(lis); err != nil {
        s.logger.Fatalln("Failed to serve:", err)
    }

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