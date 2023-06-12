package main

import (
	"context"
	"fmt"

	"github.com/gt2rz/micro-auth/internal/routes"
	"github.com/gt2rz/micro-auth/internal/servers"
)

func init() {
	fmt.Println("Initializing...")
}

func main() {

	httpServer, err := servers.NewHttpServer(context.Background())
	if err != nil {
		panic(err.Error())
	}

	httpServer.Start(routes.InitRoutes)
}
