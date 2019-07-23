package services

import (
	"context"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// AddAccount ...
func (s *Server) AddAccount(context context.Context, in *AddAccountRequest) (*AddAccountResponse, error) {

	ctx := appContext.NewContext(context, "AddAccount")
	defer ctx.Close()

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &AddAccountResponse{Id: account.ID}, nil
}
