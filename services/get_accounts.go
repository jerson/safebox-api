package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// GetAccounts ...
func (s *Server) GetAccounts(context context.Context, in *AccountsRequest) (*AccountsResponse, error) {

	ctx := appContext.NewContext(context, "GetAccounts")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	repository := repositories.NewAccountRepository(ctx)
	accounts, err := repository.ListByUserID(user.ID, 0, 1000, "id", "desc")
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	var response []*AccountSingle
	for _, account := range accounts.Items {
		response = append(response, &AccountSingle{
			Id:       account.ID,
			Hint:     account.Hint,
			Username: account.Username,
			Label:    account.Label,
		})
	}

	return &AccountsResponse{Accounts: response}, nil
}
