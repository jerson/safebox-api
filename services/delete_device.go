package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// DeleteDevice ...
func (s *Server) DeleteDevice(context context.Context, in *DeleteDeviceRequest) (*DeleteDeviceResponse, error) {
	ctx := appContext.NewContext(context, "DeleteDevice")
	defer ctx.Close()

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	repository := repositories.NewDeviceRepository(ctx)
	account, err := repository.FindOneByID(in.Id)
	if err != nil {
		return nil, err
	}

	if account.UserID != user.ID {
		return nil, errors.New("not allowed")
	}

	err = repository.Delete(*account)
	if err != nil {
		return nil, err
	}

	return &DeleteDeviceResponse{Success: true}, nil
}
