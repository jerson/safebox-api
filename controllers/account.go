package controllers

import (
	"github.com/kataras/iris"
	"net/http"
	"safebox.jerson.dev/api/forms"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

//AccountController ...
func AccountController(app *iris.Application) {

	app.Get("/account", func(c iris.Context) {

		ctx := context.NewIris(c, "Account.List")
		defer ctx.Close()

		limit, _ := ctx.GetURLParamInt("limit")
		offset, _ := ctx.GetURLParamInt("offset")
		query := ctx.GetURLParam("query")
		sort := ctx.GetURLParam("sort")
		sortType := ctx.GetURLParam("sort_type")
		repo := repositories.NewAccountRepository(ctx)

		var err error
		var result models.AccountList

		if query != "" {
			result, err = repo.SearchList(query, offset, limit)
		} else {
			result, err = repo.List(offset, limit, sort, sortType)
		}
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}

		ctx.SendResponse(http.StatusOK, result)

	})

	app.Get("/account/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Account.Get")
			defer ctx.Close()
			id, _ := ctx.GetParamInt64("id")
			repo := repositories.NewAccountRepository(ctx)
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			ctx.SendResponse(http.StatusOK, item)

		})

	app.Post("/account", func(c iris.Context) {

		ctx := context.NewIris(c, "Account.Post")
		defer ctx.Close()

		formData := forms.AccountForm{}
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
		item := models.Account{}
		itemBinded := ctx.BindRequest(item, formData)

		repo := repositories.NewAccountRepository(ctx)
		result, err := repo.Create(itemBinded.(models.Account))
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}
		ctx.SendResponse(http.StatusOK, result)

	})

	app.Put("/account/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Account.Put")
			defer ctx.Close()

			repo := repositories.NewAccountRepository(ctx)
			id, _ := ctx.GetParamInt64("id")
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			formData := forms.AccountForm{}
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

	app.Delete("/account/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Account.Delete")
			defer ctx.Close()

			repo := repositories.NewAccountRepository(ctx)
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
