package controllers

import (
	"github.com/kataras/iris"
	"net/http"
	"safebox.jerson.dev/api/forms"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

//UserController ...
func UserController(app *iris.Application) {

	app.Get("/user", func(c iris.Context) {

		ctx := context.NewIris(c, "User.List")
		defer ctx.Close()

		limit, _ := ctx.GetURLParamInt("limit")
		offset, _ := ctx.GetURLParamInt("offset")
		query := ctx.GetURLParam("query")
		sort := ctx.GetURLParam("sort")
		sortType := ctx.GetURLParam("sort_type")
		repo := repositories.NewUserRepository(ctx)

		var err error
		var result models.UserList

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

	app.Get("/user/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "User.Get")
			defer ctx.Close()
			id, _ := ctx.GetParamInt64("id")
			repo := repositories.NewUserRepository(ctx)
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			ctx.SendResponse(http.StatusOK, item)

		})

	app.Post("/user", func(c iris.Context) {

		ctx := context.NewIris(c, "User.Post")
		defer ctx.Close()

		formData := forms.UserForm{}
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
		item := models.User{}
		itemBinded := ctx.BindRequest(item, formData)

		repo := repositories.NewUserRepository(ctx)
		result, err := repo.Create(itemBinded.(models.User))
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}
		ctx.SendResponse(http.StatusOK, result)

	})

	app.Put("/user/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "User.Put")
			defer ctx.Close()

			repo := repositories.NewUserRepository(ctx)
			id, _ := ctx.GetParamInt64("id")
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			formData := forms.UserForm{}
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

	app.Delete("/user/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "User.Delete")
			defer ctx.Close()

			repo := repositories.NewUserRepository(ctx)
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
