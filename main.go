package main

import (
	"fmt"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/metrics"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
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

	server := fmt.Sprintf(":%d", config.Vars.Server.Port)
	fmt.Println("running: ", server)
	service := rpc.NewHTTPService()
	service.AddFunction("Ping", func() string {
		return time.Now().String()
	})

	service.AddFunction("GetStatus", func() (string, error) {

		ctx := context.NewSingle("command")
		defer ctx.Close()

		return "tes", nil
	})

	metrics.RPC(ctx, service)

	http.Handle("/", service)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		panic(err)
	}
}
