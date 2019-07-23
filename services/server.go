package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
)

// Server ...
type Server struct{}

// Ping ...
func (s *Server) Ping(context context.Context, in *PingRequest) (*PingResponse, error) {

	ctx := appContext.NewContext(context, "Ping")
	defer ctx.Close()

	log := ctx.GetLogger("main")
	log.Info("Ping")

	return &PingResponse{Name: "Hello Ping"}, nil
}
