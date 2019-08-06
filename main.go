package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/services"

	"github.com/soheilhy/cmux"
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
		&models.UserLocation{},
		&models.Product{},
		&models.Purchase{},
		&models.User{},
		&models.Device{},
	)
}

func main() {

	ctx := context.NewContextSingle("main")
	defer ctx.Close()

	migrate(ctx)

	port := fmt.Sprintf(":%d", config.Vars.Server.Port)
	fmt.Println("running: ", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	services.RegisterServicesServer(server, &services.Server{})

	mux := cmux.New(listener)
	grpcListener := mux.Match(cmux.HTTP2())
	grpcWebListener := mux.Match(cmux.HTTP1HeaderFieldPrefix("content-type", "application/grpc-web"))
	httpListener := mux.Match(cmux.HTTP1())

	group := new(errgroup.Group)
	group.Go(func() error { return grpcServe(server, grpcListener) })
	group.Go(func() error { return grpcWebServe(server, grpcWebListener) })
	group.Go(func() error { return httpServe(httpListener) })
	group.Go(func() error { return mux.Serve() })

	err = group.Wait()
	if err != nil {
		panic(err)
	}
}

func httpServe(listen net.Listener) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("hi"))
	})

	httpServer := &http.Server{Handler: mux}

	return httpServer.Serve(listen)
}
func grpcWebServe(server *grpc.Server, listen net.Listener) error {

	wrappedGRPC := grpcweb.WrapServer(server)
	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGRPC.IsGrpcWebRequest(req) {
			wrappedGRPC.ServeHTTP(resp, req)
			return
		}
		http.DefaultServeMux.ServeHTTP(resp, req)
	})

	httpServer := &http.Server{Handler: handler}
	return httpServer.Serve(listen)
}

func grpcServe(server *grpc.Server, listen net.Listener) error {
	return server.Serve(listen)
}
