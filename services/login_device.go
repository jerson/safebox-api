package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
)

// LoginWithDevice ...
func (s *Server) LoginWithDevice(context context.Context, in *LoginDeviceRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "LoginWithDevice")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByDevice(ctx, in.PublicKey)
	if err != nil {
		log.Error(err)
		return nil, errors.New("the device is not registered")
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return response, nil
}
