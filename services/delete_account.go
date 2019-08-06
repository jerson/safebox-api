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

	err = repository.Delete(*account)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &DeleteAccountResponse{}, nil
}
