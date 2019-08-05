package services

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// SendLocation ...
func (s *Server) SendLocation(context context.Context, in *SendLocationRequest) (*SendLocationResponse, error) {

	ctx := appContext.NewContext(context, "SendLocation")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	if !user.LocationEnabled {
		return nil, errors.New("location not enabled")
	}

	repository := repositories.NewUserLocationRepository(ctx)
	location, err := repository.FindOneByUserID(user.ID)
	if err != nil {
		locationInput := models.UserLocation{
			UserID:    user.ID,
			Latitude:  in.Latitude,
			Longitude: in.Longitude,
		}
		_, err = repository.Create(locationInput)

	} else {
		err = repository.UpdateColumns(*location, map[string]string{
			"latitude":  in.Latitude,
			"longitude": in.Longitude,
		})
	}
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &SendLocationResponse{}, nil
}
