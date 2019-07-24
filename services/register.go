package services

import (
	"context"
	"errors"
	"fmt"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// Register ...
func (s *Server) Register(context context.Context, in *RegisterRequest) (*RegisterResponse, error) {

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

	fmt.Println("register1")

	user, err := repository.Create(userInput)
	if err != nil {
		return nil, err
	}
	fmt.Println("register2")

	accessToken, err := getAccessToken(ctx, *user)
	if err != nil {
		return nil, err
	}
	fmt.Println("register3")

	return &RegisterResponse{AccessToken: accessToken}, nil
}
