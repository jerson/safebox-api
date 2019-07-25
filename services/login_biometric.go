package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/openpgp"
	"safebox.jerson.dev/api/repositories"
)

// LoginBiometric ...
func (s *Server) LoginBiometric(context context.Context, in *LoginBiometricRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "Login")
	defer ctx.Close()

	repository := repositories.NewUserRepository(ctx)

	user, err := repository.FindOneByUsername(in.Username)
	if err != nil {
		return nil, err
	}

	pgp := openpgp.NewOpenPGP()
	// FIXME aqui debe implementarse biometric publicKey
	_, err = pgp.ReadPrivateKey(user.PrivateKey, in.Username)
	if err != nil {
		return nil, err
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
