package main

import (
	"log"

	"github.com/minhquy1903/online-test/api-gateway/pkg/auth"
	"github.com/minhquy1903/online-test/api-gateway/pkg/config"
)

func main() {
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    authSvc := *auth.RegisterRoutes(r, &c)

    r.Run(c.Port)
}