package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// DeleteAccount ...
func (s *Server) DeleteAccount(context context.Context, in *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	ctx := appContext.NewContext(context, "DeleteAccount")
	defer ctx.Close()

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	repository := repositories.NewAccountRepository(ctx)
	account, err := repository.FindOneByID(in.Id)
	if err != nil {
		return nil, err
	}

	if account.UserID != user.ID {
		return nil, errors.New("not allowed")
	}

	err = repository.Delete(*account)
	if err != nil {
		return nil, err
	}

	return &DeleteAccountResponse{Success: true}, nil
}
