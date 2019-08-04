package services

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/util"
	"safebox.jerson.dev/api/repositories"
)

// AddDevice ...
func (s *Server) AddDevice(context context.Context, in *AddDeviceRequest) (*AddDeviceResponse, error) {

	ctx := appContext.NewContext(context, "AddDevice")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	hash := util.SHA512(in.PublicKey)
	repository := repositories.NewDeviceRepository(ctx)
	deviceInput := models.Device{
		UserID:    user.ID,
		PublicKey: in.PublicKey,
		Name:      in.Name,
		UID:       in.Uid,
		Hash:      hash,
	}

	device, err := repository.Create(deviceInput)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &AddDeviceResponse{Id: device.ID}, nil
}
