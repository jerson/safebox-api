package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/openpgp"
	"safebox.jerson.dev/api/repositories"
)

// Login ...
func (s *Server) Login(context context.Context, in *LoginRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "Login")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	repository := repositories.NewUserRepository(ctx)

	user, err := repository.FindOneByUsername(in.Username)
	if err != nil {
		log.Error(err)
		return nil, errors.New("invalid credentials")
	}

	pgp := openpgp.NewOpenPGP()
	_, err = pgp.ReadPrivateKey(user.PrivateKey, in.Password)
	if err != nil {
		log.Error(err)
		return nil, errors.New("invalid credentials")
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return response, nil
}
