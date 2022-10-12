package server

import (
	"fmt"
	"net"

	"github.com/minhquy1903/online-test/auth-service/config"
	"github.com/minhquy1903/online-test/auth-service/pb"
	"github.com/minhquy1903/online-test/auth-service/service"
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

func (s *Server) Run() {
	h := service.NewHandler(s.db, s.cfg)

	lis, err := net.Listen("tcp", ":"+s.cfg.Server.Port)

	if err != nil {
        s.logger.Fatalln("Failed to listing:", err)
    }

    grpcServer := grpc.NewServer()

    pb.RegisterAuthServiceServer(grpcServer, h)
	
	fmt.Println("Service is listening on port:", s.cfg.Server.Port)

	
    if err := grpcServer.Serve(lis); err != nil {
		s.logger.Fatalln("Failed to serve:", err)
	} 

}