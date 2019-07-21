package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"net"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	pb "safebox.jerson.dev/api/services"
)

func init() {
	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}
}
func migrate(ctx context.Base) {
	cn, err := db.Setup(ctx)
	if err != nil {
		panic(err)
	}
	cn.AutoMigrate(
		&models.Account{},
		&models.AuditLog{},
		&models.Product{},
		&models.Purchase{},
		&models.User{},
	)
}

func main() {

	ctx := context.NewSingle("main")
	defer ctx.Close()

	migrate(ctx)

	port := fmt.Sprintf(":%d", config.Vars.Server.Port)
	fmt.Println("running: ", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterServicesServer(server, &pb.Server{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
