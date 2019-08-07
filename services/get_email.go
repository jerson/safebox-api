package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
)

// GetEmail ...
func (s *Server) GetEmail(context context.Context, in *GetEmailRequest) (*GetEmailResponse, error) {

	ctx := appContext.NewContext(context, "GetEmail")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	return &GetEmailResponse{Email: user.Email}, nil
}
