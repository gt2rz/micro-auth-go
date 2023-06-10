package main

import (
	"context"
	"fmt"

	"github.com/gt2rz/micro-auth/internal/routes"
	"github.com/gt2rz/micro-auth/internal/servers"
	"github.com/gt2rz/micro-auth/internal/utils"
)

func init() {
	fmt.Println("Initializing...")
	utils.LoadEnvs()
}

func main() {

	httpServer, err := servers.NewHttpServer(context.Background())
	if err != nil {
		panic(err.Error())
	}

	httpServer.Start(routes.InitRoutes)
}
