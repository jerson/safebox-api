package services

import (
	"context"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/util"
	"safebox.jerson.dev/api/repositories"
)

// AddDevice ...
func (s *Server) AddDevice(context context.Context, in *AddDeviceRequest) (*AddDeviceResponse, error) {

	ctx := appContext.NewContext(context, "AddDevice")
	defer ctx.Close()

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	hash := util.SHA512(in.PublicKey)
	repository := repositories.NewDeviceRepository(ctx)
	deviceInput := models.Device{
		UserID:    user.ID,
		PublicKey: in.PublicKey,
		Name:      in.Name,
		Hash:      hash,
	}

	device, err := repository.Create(deviceInput)
	if err != nil {
		return nil, err
	}

	return &AddDeviceResponse{Id: device.ID}, nil
}
