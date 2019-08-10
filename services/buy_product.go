package services

import (
	"context"
	"errors"
	"github.com/awa/go-iap/appstore"
	"github.com/awa/go-iap/playstore"
	"io/ioutil"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/config"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
	"time"
)

// BuyProduct ...
func (s *Server) BuyProduct(contextApp context.Context, in *BuyProductRequest) (*BuyProductResponse, error) {

	ctx := appContext.NewContext(contextApp, "BuyProduct")
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

		jsonKey, err := ioutil.ReadFile(config.Vars.Purchase.GooglePlayFile)
		if err != nil {
			log.Error(err)
			return nil, errors.New("error verifying payment")
		}

		client, err := playstore.New(jsonKey)
		if err != nil {
			log.Error(err)
			return nil, errors.New("error verifying payment")
		}

		ctx := context.Background()
		_, err = client.VerifyProduct(ctx, config.Vars.Purchase.PackageID, product.Slug, in.Payload)
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

		ctx := context.Background()
		err := client.Verify(ctx, req, resp)
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
