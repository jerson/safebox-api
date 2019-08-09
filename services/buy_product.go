package services

import (
	"context"
	"errors"
	"github.com/awa/go-iap/appstore"
	"github.com/awa/go-iap/playstore"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/config"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
	"time"
)

// BuyProduct ...
func (s *Server) BuyProduct(context context.Context, in *BuyProductRequest) (*BuyProductResponse, error) {

	ctx := appContext.NewContext(context, "BuyProduct")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	productRepository := repositories.NewProductRepository(ctx)
	product, err := productRepository.FindOneBySlug(in.Slug)
	if err != nil {
		log.Error(err)
		return nil, errors.New("product not found")
	}

	if in.Type == "android" {

		client, err := playstore.New([]byte(in.Payload))
		if err != nil {
			log.Error(err)
			return nil, errors.New("error verifying payment")
		}
		_, err = client.VerifyProduct(context, config.Vars.Payment.PackageID, product.Slug, in.Token)
		if err != nil {
			log.Error(err)
			return nil, errors.New("error verifying payment")
		}

	} else if in.Type == "ios" {
		client := appstore.New()
		req := appstore.IAPRequest{
			ReceiptData: in.Payload,
		}
		resp := &appstore.IAPResponse{}
		err := client.Verify(context, req, resp)
		if err != nil {
			log.Error(err)
			return nil, errors.New("error verifying payment")
		}
	} else {
		return nil, errors.New("type not supported")
	}

	purchaseInput := models.Purchase{
		UserID:    user.ID,
		ProductID: product.ID,
		Payload:   in.Payload,
	}
	repository := repositories.NewPurchaseRepository(ctx)
	purchase, err := repository.Create(purchaseInput)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &BuyProductResponse{
		Id:   purchase.ID,
		Date: purchase.Date.Format(time.RFC3339),
	}, nil
}
