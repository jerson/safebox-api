package services

import (
	"context"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

// GetAccounts ...
func (s *Server) GetAccounts(context context.Context, in *AccountsRequest) (*AccountsResponse, error) {

	ctx := appContext.NewContext(context, "GetAccounts")
	defer ctx.Close()

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	repository := repositories.NewAccountRepository(ctx)
	accounts, err := repository.ListByUserID(user.ID, 0, 100, "id", "desc")
	if err != nil {
		return nil, err
	}

	var response []*AccountSingle
	for _, account := range accounts.Items {
		response = append(response, &AccountSingle{
			Hint:     account.Hint,
			Username: account.Username,
			Label:    account.Label,
		})
	}

	return &AccountsResponse{Accounts: response}, nil
}
