package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Gemba-Kaizen/menu-service/config"
	"github.com/Gemba-Kaizen/menu-service/internal/db"
	repository "github.com/Gemba-Kaizen/menu-service/internal/repository/menu"
	api "github.com/Gemba-Kaizen/menu-service/pkg/api/menu"
	"github.com/Gemba-Kaizen/menu-service/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	// Init DB
	h := db.Init(c.DBUrl)

	menuRepo := &repository.MenuRepository{H: &h}

	// Init handlers
	menuHandler := &api.MenuHandler{MenuRepo: menuRepo}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed at listen: ", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	grpcService := grpc.NewServer()

	// Register each handler endpoint to grpc Server
	pb.RegisterMenuServiceServer(grpcService, menuHandler)
	// pb.RegisterService2ServiceServer(grpcServer, service2Handler)

	if err := grpcService.Serve(lis); err != nil {
		log.Fatalln("Failed at serve: ", err)
	}
}
