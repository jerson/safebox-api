package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// RefreshToken ...
func (s *Server) RefreshToken(context context.Context, in *RefreshTokenRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "RefreshToken")
	defer ctx.Close()

	repository := repositories.NewAccessTokenRepository(ctx)
	accessToken, err := repository.FindOneByToken(in.AccessToken)
	if err != nil {
		return nil, err
	}

	userID := accessToken.UserID
	err = repository.Delete(*accessToken)
	if err != nil {
		return nil, err
	}

	userRepo := repositories.NewUserRepository(ctx)
	user, err := userRepo.FindOneByID(userID)
	if err != nil {
		return nil, err
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
