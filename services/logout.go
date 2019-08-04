package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// Logout ...
func (s *Server) Logout(context context.Context, in *LogoutRequest) (*LogoutResponse, error) {

	ctx := appContext.NewContext(context, "Logout")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	repository := repositories.NewAccessTokenRepository(ctx)
	accessToken, err := repository.FindOneByToken(in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}
	err = repository.Delete(*accessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}
	return &LogoutResponse{}, nil
}
