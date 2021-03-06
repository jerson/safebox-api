package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// GetAccount ...
func (s *Server) GetAccount(context context.Context, in *AccountRequest) (*AccountResponse, error) {
	ctx := appContext.NewContext(context, "GetAccount")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	repository := repositories.NewAccountRepository(ctx)
	account, err := repository.FindOneByID(in.Id)
	if err != nil {
		log.Error(err)
		return nil, errors.New("account has already been deleted")
	}

	if account.UserID != user.ID {
		return nil, errors.New("not allowed")
	}

	response := &Account{
		Hint:     account.Hint,
		Password: account.Password,
		Username: account.Username,
		Label:    account.Label,
	}
	return &AccountResponse{Account: response}, nil
}
