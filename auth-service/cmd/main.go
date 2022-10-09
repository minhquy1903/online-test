package main

import (
	"fmt"
	"log"

	"github.com/minhquy1903/online-test/auth-service/config"
	"github.com/minhquy1903/online-test/auth-service/internal/server"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig(".")

	if err != nil{
		fmt.Println("fail")
		log.Fatalf("LoadConfig: %v", err)
	}

	// db := db.GetPostgresInstance(cfg)

	s := server.NewServer(cfg, nil, logrus.New(), nil)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

