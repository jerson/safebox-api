package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// DisableLocation ...
func (s *Server) DisableLocation(context context.Context, in *DisableLocationRequest) (*DisableLocationResponse, error) {

	ctx := appContext.NewContext(context, "DisableLocation")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	repository := repositories.NewUserRepository(ctx)
	err = repository.UpdateColumns(*user, map[string]string{
		"email":            "",
		"location_enabled": "0",
	})
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &DisableLocationResponse{}, nil
}
