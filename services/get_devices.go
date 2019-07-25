package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
	"time"
)

// GetDevices ...
func (s *Server) GetDevices(context context.Context, in *DevicesRequest) (*DevicesResponse, error) {

	ctx := appContext.NewContext(context, "GetDevices")
	defer ctx.Close()

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	repository := repositories.NewDeviceRepository(ctx)
	accounts, err := repository.ListByUserID(user.ID, 0, 1000, "id", "desc")
	if err != nil {
		return nil, err
	}

	var response []*Device
	for _, device := range accounts.Items {
		response = append(response, &Device{
			Name:        device.Name,
			Id:          device.ID,
			DateCreated: device.DateCreated.Format(time.RFC3339),
		})
	}

	return &DevicesResponse{Devices: response}, nil
}
