package services

import (
	"errors"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/util"
	"safebox.jerson.dev/api/repositories"
	"time"
)

func getUserByDevice(ctx context.Context, publicKey string) (*models.User, error) {

	hash := util.SHA512(publicKey)
	repository := repositories.NewDeviceRepository(ctx)
	device, err := repository.FindOneByHash(hash)
	if err != nil {
		return nil, err
	}

	userRepository := repositories.NewUserRepository(ctx)
	user, err := userRepository.FindOneByID(device.UserID)
	if err != nil {
		return nil, err
	}

	ctx.SetUser(user.ID)

	return user, nil
}

func getUserByToken(ctx context.Context, token string) (*models.User, error) {

	repository := repositories.NewAccessTokenRepository(ctx)
	accessToken, err := repository.FindOneByToken(token)
	if err != nil {
		return nil, err
	}

	diff := accessToken.DateExpire.Sub(time.Now())
	if diff < time.Second*0 {
		return nil, errors.New("expired session, login again please")
	}

	userRepository := repositories.NewUserRepository(ctx)
	user, err := userRepository.FindOneByID(accessToken.UserID)
	if err != nil {
		return nil, err
	}

	ctx.SetUser(user.ID)

	return user, nil
}

func getAuthResponse(ctx context.Context, user models.User) (*AuthResponse, error) {

	repository := repositories.NewAccessTokenRepository(ctx)
	dateExpire := time.Now().Add(time.Minute * 5)

	randomToken, err := util.GenerateRandomASCIIString(128)
	if err != nil {
		return nil, err
	}
	accessTokenInput := models.AccessToken{
		UserID:     user.ID,
		DateExpire: &dateExpire,
		Token:      randomToken,
	}
	accessToken, err := repository.Create(accessTokenInput)
	if err != nil {
		return nil, err
	}

	userRepo := repositories.NewUserRepository(ctx)
	err = userRepo.UpdateSingle(user, "date_connected", time.Now())
	if err != nil {
		return nil, err
	}

	ctx.SetUser(user.ID)

	return &AuthResponse{
		AccessToken: accessToken.Token,
		DateExpire:  accessToken.DateExpire.Format(time.RFC3339),
		Date:  time.Now().Format(time.RFC3339),
		KeyPair: &KeyPairResponse{
			PrivateKey: user.PrivateKey,
			PublicKey:  user.PublicKey,
		},
	}, nil
}
