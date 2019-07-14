package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"net/http"
	"safebox.jerson.dev/api/controllers"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
)

func init() {
	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}
}

func main() {

	ctx := context.NewSingle("api")
	defer ctx.Close()

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

	app := iris.New()
	app.Use(recover.New())
	app.Use(cors.Default())

	controllers.MetricsController(app)
	controllers.BaseController(app)
	controllers.AccountController(app)
	controllers.AuditLogController(app)
	controllers.ProductController(app)
	controllers.PurchaseController(app)
	controllers.UserController(app)

	log := ctx.GetLogger("main")
	port := fmt.Sprintf(":%d", config.Vars.Server.Port)
	log.Infof("running %s", port)

	app.Build()
	err = gracehttp.Serve(
		&http.Server{Addr: port, Handler: app},
	)
	if err != nil {
		panic(err)
	}

}
