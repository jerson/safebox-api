package services

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// Register ...
func (s *Server) Register(context context.Context, in *RegisterRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "Register")
	defer ctx.Close()

	repository := repositories.NewUserRepository(ctx)

	user, _ := repository.FindOneByUsername(in.Username)
	if user != nil {
		return nil, errors.New("already registered")
	}

	userInput := models.User{
		PrivateKey: in.PrivateKey,
		PublicKey:  in.PublicKey,
		Username:   in.Username,
	}

	user, err := repository.Create(userInput)
	if err != nil {
		return nil, err
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
