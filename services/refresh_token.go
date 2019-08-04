package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// RefreshToken ...
func (s *Server) RefreshToken(context context.Context, in *RefreshTokenRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "RefreshToken")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	repository := repositories.NewAccessTokenRepository(ctx)
	accessToken, err := repository.FindOneByToken(in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	userID := accessToken.UserID
	err = repository.Delete(*accessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	userRepo := repositories.NewUserRepository(ctx)
	user, err := userRepo.FindOneByID(userID)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return response, nil
}
