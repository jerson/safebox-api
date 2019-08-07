package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// EnableLocation ...
func (s *Server) EnableLocation(context context.Context, in *EnableLocationRequest) (*EnableLocationResponse, error) {

	ctx := appContext.NewContext(context, "EnableLocation")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}
	if in.Email == "" {
		return nil, errors.New("email required")
	}

	repository := repositories.NewUserRepository(ctx)
	err = repository.UpdateColumns(*user, map[string]string{
		"email":            in.Email,
		"location_enabled": "1",
	})
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &EnableLocationResponse{}, nil
}
