package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// Logout ...
func (s *Server) Logout(context context.Context, in *LogoutRequest) (*LogoutResponse, error) {

	ctx := appContext.NewContext(context, "Logout")
	defer ctx.Close()

	repository := repositories.NewAccessTokenRepository(ctx)
	accessToken, err := repository.FindOneByToken(in.AccessToken)
	if err != nil {
		return nil, err
	}
	err = repository.Delete(*accessToken)
	if err != nil {
		return nil, err
	}
	return &LogoutResponse{}, nil
}
