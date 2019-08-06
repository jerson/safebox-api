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

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	repository := repositories.NewDeviceRepository(ctx)
	account, err := repository.FindOneByID(in.Id)
	if err != nil {
		log.Error(err)
		return nil, errors.New("device has already been deleted")
	}

	if account.UserID != user.ID {
		return nil, errors.New("not allowed")
	}

	err = repository.Delete(*account)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &DeleteDeviceResponse{}, nil
}
