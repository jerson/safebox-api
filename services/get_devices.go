package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
	"time"
)

// GetDevices ...
func (s *Server) GetDevices(context context.Context, in *DevicesRequest) (*DevicesResponse, error) {

	ctx := appContext.NewContext(context, "GetDevices")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	repository := repositories.NewDeviceRepository(ctx)
	accounts, err := repository.ListByUserID(user.ID, 0, 1000, "id", "desc")
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	var response []*Device
	for _, device := range accounts.Items {
		response = append(response, &Device{
			Name:        device.Name,
			Id:          device.ID,
			Uid:         device.UID,
			DateCreated: device.DateCreated.Format(time.RFC3339),
		})
	}

	return &DevicesResponse{Devices: response}, nil
}
