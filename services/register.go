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

	log := ctx.GetLogger("RPC")

	repository := repositories.NewUserRepository(ctx)

	userFound, err := repository.FindOneByUsername(in.Username)
	if userFound != nil {
		log.Error(err)
		return nil, errors.New("already registered")
	}

	userInput := models.User{
		PrivateKey: in.PrivateKey,
		PublicKey:  in.PublicKey,
		Username:   in.Username,
	}

	user, err := repository.Create(userInput)
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
