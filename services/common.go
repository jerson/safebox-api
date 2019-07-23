package services

import (
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

	accessTokenInput := models.AccessToken{
		UserID:     user.ID,
		DateExpire: &dateExpire,
		Token:      util.UniqueID(),
	}
	accessToken, err := repository.Create(accessTokenInput)
	if err != nil {
		return token, err
	}

	ctx.SetUser(user.ID)

	return accessToken.Token, nil
}
