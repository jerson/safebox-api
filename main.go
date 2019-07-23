package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"net/http"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/services"
	"time"
)

func init() {
	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}
}
func migrate(ctx context.Context) {
	cn, err := db.Setup(ctx)
	if err != nil {
		panic(err)
	}
	cn.AutoMigrate(
		&models.AccessToken{},
		&models.Account{},
		&models.AuditLog{},
		&models.Product{},
		&models.Purchase{},
		&models.User{},
	)
}

func main() {

	ctx := context.NewContextSingle("main")
	defer ctx.Close()

	migrate(ctx)

	port := fmt.Sprintf(":%d", config.Vars.Server.Port)
	fmt.Println("running: ", port)

	server := grpc.NewServer()
	services.RegisterServicesServer(server, &services.Server{})

	wrappedGRPC := grpcweb.WrapServer(server)
	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGRPC.IsGrpcWebRequest(req) {
			wrappedGRPC.ServeHTTP(resp, req)
			return
		}
		http.DefaultServeMux.ServeHTTP(resp, req)
	})

	s := &http.Server{
		Addr:           port,
		Handler:        handler,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
