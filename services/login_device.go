package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
)

// LoginWithDevice ...
func (s *Server) LoginWithDevice(context context.Context, in *LoginDeviceRequest) (*AuthResponse, error) {

	ctx := appContext.NewContext(context, "LoginWithDevice")
	defer ctx.Close()

	user, err := getUserByDevice(ctx, in.PublicKey)
	if err != nil {
		return nil, err
	}

	response, err := getAuthResponse(ctx, *user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
