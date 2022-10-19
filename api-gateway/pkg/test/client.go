package auth

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/online-test/api-gateway/pkg/auth/pb"
	"github.com/minhquy1903/online-test/api-gateway/pkg/auth/routes"
	"github.com/minhquy1903/online-test/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
    Client pb.AuthServiceClient
}

func (svc *ServiceClient) Register(ctx echo.Context) error {
    routes.Register(ctx, svc.Client)
    return nil
}

func (svc *ServiceClient) Login(ctx echo.Context) error {
    routes.Login(ctx, svc.Client)
    return nil
}


func InitServiceClient(c *config.Config) pb.AuthServiceClient {
    // using WithInsecure() because no SSL running
    
    cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    return pb.NewAuthServiceClient(cc)
}