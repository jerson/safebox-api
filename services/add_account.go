package services

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// AddAccount ...
func (s *Server) AddAccount(context context.Context, in *AddAccountRequest) (*AddAccountResponse, error) {

	ctx := appContext.NewContext(context, "AddAccount")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	repository := repositories.NewAccountRepository(ctx)
	accountInput := models.Account{
		UserID:   user.ID,
		Username: in.Account.Username,
		Password: in.Account.Password,
		Hint:     in.Account.Hint,
		Label:    in.Account.Label,
	}

	account, err := repository.Create(accountInput)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &AddAccountResponse{Id: account.ID}, nil
}
