package controllers

import (
	"github.com/kataras/iris"
	"net/http"
	"safebox.jerson.dev/api/forms"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

//PurchaseController ...
func PurchaseController(app *iris.Application) {

	app.Get("/purchase", func(c iris.Context) {

		ctx := context.NewIris(c, "Purchase.List")
		defer ctx.Close()

		limit, _ := ctx.GetURLParamInt("limit")
		offset, _ := ctx.GetURLParamInt("offset")
		sort := ctx.GetURLParam("sort")
		sortType := ctx.GetURLParam("sort_type")
		repo := repositories.NewPurchaseRepository(ctx)

		var err error
		var result models.PurchaseList

		result, err = repo.List(offset, limit, sort, sortType)
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}

		ctx.SendResponse(http.StatusOK, result)

	})

	app.Get("/purchase/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Purchase.Get")
			defer ctx.Close()
			id, _ := ctx.GetParamInt64("id")
			repo := repositories.NewPurchaseRepository(ctx)
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			ctx.SendResponse(http.StatusOK, item)

		})

	app.Post("/purchase", func(c iris.Context) {

		ctx := context.NewIris(c, "Purchase.Post")
		defer ctx.Close()

		formData := forms.PurchaseForm{}
		err := ctx.ReadForm(&formData)
		if err != nil {
			ctx.SendError(http.StatusBadRequest, err)
			return
		}
		err = formData.IsValid()
		if err != nil {
			ctx.SendError(http.StatusBadRequest, err)
			return
		}
		item := models.Purchase{}
		itemBinded := ctx.BindRequest(item, formData)

		repo := repositories.NewPurchaseRepository(ctx)
		result, err := repo.Create(itemBinded.(models.Purchase))
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}
		ctx.SendResponse(http.StatusOK, result)

	})

	app.Put("/purchase/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Purchase.Put")
			defer ctx.Close()

			repo := repositories.NewPurchaseRepository(ctx)
			id, _ := ctx.GetParamInt64("id")
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			formData := forms.PurchaseForm{}
			err = ctx.ReadForm(&formData)
			if err != nil {
				ctx.SendError(http.StatusBadRequest, err)
				return
			}
			err = formData.IsValid()
			if err != nil {
				ctx.SendError(http.StatusBadRequest, err)
				return
			}
			itemBinded := ctx.BindRequest(item, formData)

			err = repo.Update(*item, itemBinded)
			if err != nil {
				ctx.SendError(http.StatusInternalServerError, err)
				return
			}
			ctx.SendResponse(http.StatusOK, item)

		})

	app.Delete("/purchase/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Purchase.Delete")
			defer ctx.Close()

			repo := repositories.NewPurchaseRepository(ctx)
			id, _ := ctx.GetParamInt64("id")
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			err = repo.Delete(*item)
			if err != nil {
				ctx.SendError(http.StatusInternalServerError, err)
				return
			}

			ctx.SendResponse(http.StatusOK, item)

		})
}
