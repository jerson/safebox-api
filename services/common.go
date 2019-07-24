package services

import (
	"errors"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/util"
	"safebox.jerson.dev/api/repositories"
	"time"
)

func getUserByToken(ctx context.Context, token string) (*models.User, error) {

	repository := repositories.NewAccessTokenRepository(ctx)
	accessToken, err := repository.FindOneByToken(token)
	if err != nil {
		return nil, err
	}

	diff := accessToken.DateExpire.Sub(time.Now())
	if diff < time.Second*0 {
		return nil, errors.New("expired token")
	}

	userRepository := repositories.NewUserRepository(ctx)
	user, err := userRepository.FindOneByID(accessToken.UserID)
	if err != nil {
		return nil, err
	}

	ctx.SetUser(user.ID)

	return user, nil
}

func getAccessToken(ctx context.Context, user models.User) (string, error) {

	var token string

	repository := repositories.NewAccessTokenRepository(ctx)
	dateExpire := time.Now().Add(time.Minute * 5)

	randomToken, err := util.GenerateRandomASCIIString(128)
	if err != nil {
		return token, err
	}
	accessTokenInput := models.AccessToken{
		UserID:     user.ID,
		DateExpire: &dateExpire,
		Token:      randomToken,
	}
	accessToken, err := repository.Create(accessTokenInput)
	if err != nil {
		return token, err
	}

	userRepo := repositories.NewUserRepository(ctx)
	err = userRepo.UpdateSingle(user, "date_connected", time.Now())
	if err != nil {
		return token, err
	}

	ctx.SetUser(user.ID)

	return accessToken.Token, nil
}
