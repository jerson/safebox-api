package controllers

import (
	"github.com/kataras/iris"
	"net/http"
	"safebox.jerson.dev/api/forms"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

//ProductController ...
func ProductController(app *iris.Application) {

	app.Get("/product", func(c iris.Context) {

		ctx := context.NewIris(c, "Product.List")
		defer ctx.Close()

		limit, _ := ctx.GetURLParamInt("limit")
		offset, _ := ctx.GetURLParamInt("offset")
		query := ctx.GetURLParam("query")
		sort := ctx.GetURLParam("sort")
		sortType := ctx.GetURLParam("sort_type")
		repo := repositories.NewProductRepository(ctx)

		var err error
		var result models.ProductList

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

	app.Get("/product/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Product.Get")
			defer ctx.Close()
			id, _ := ctx.GetParamInt64("id")
			repo := repositories.NewProductRepository(ctx)
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			ctx.SendResponse(http.StatusOK, item)

		})

	app.Post("/product", func(c iris.Context) {

		ctx := context.NewIris(c, "Product.Post")
		defer ctx.Close()

		formData := forms.ProductForm{}
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
		item := models.Product{}
		itemBinded := ctx.BindRequest(item, formData)

		repo := repositories.NewProductRepository(ctx)
		result, err := repo.Create(itemBinded.(models.Product))
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}
		ctx.SendResponse(http.StatusOK, result)

	})

	app.Put("/product/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Product.Put")
			defer ctx.Close()

			repo := repositories.NewProductRepository(ctx)
			id, _ := ctx.GetParamInt64("id")
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			formData := forms.ProductForm{}
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

	app.Delete("/product/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "Product.Delete")
			defer ctx.Close()

			repo := repositories.NewProductRepository(ctx)
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
