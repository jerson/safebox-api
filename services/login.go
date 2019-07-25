package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/openpgp"
	"safebox.jerson.dev/api/repositories"
)

// Login ...
func (s *Server) Login(context context.Context, in *LoginRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "Login")
	defer ctx.Close()

	repository := repositories.NewUserRepository(ctx)

	user, err := repository.FindOneByUsername(in.Username)
	if err != nil {
		return nil, err
	}

	pgp := openpgp.NewOpenPGP()
	_, err = pgp.ReadPrivateKey(user.PrivateKey, in.Password)
	if err != nil {
		return nil, err
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
