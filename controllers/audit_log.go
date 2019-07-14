package controllers

import (
	"github.com/kataras/iris"
	"net/http"
	"safebox.jerson.dev/api/forms"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
)

//AuditLogController ...
func AuditLogController(app *iris.Application) {

	app.Get("/audit-log", func(c iris.Context) {

		ctx := context.NewIris(c, "AuditLog.List")
		defer ctx.Close()

		limit, _ := ctx.GetURLParamInt("limit")
		offset, _ := ctx.GetURLParamInt("offset")
		sort := ctx.GetURLParam("sort")
		sortType := ctx.GetURLParam("sort_type")
		repo := repositories.NewAuditLogRepository(ctx)

		var err error
		var result models.AuditLogList

		result, err = repo.List(offset, limit, sort, sortType)
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}

		ctx.SendResponse(http.StatusOK, result)

	})

	app.Get("/audit-log/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "AuditLog.Get")
			defer ctx.Close()
			id, _ := ctx.GetParamInt64("id")
			repo := repositories.NewAuditLogRepository(ctx)
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			ctx.SendResponse(http.StatusOK, item)

		})

	app.Post("/audit-log", func(c iris.Context) {

		ctx := context.NewIris(c, "AuditLog.Post")
		defer ctx.Close()

		formData := forms.AuditLogForm{}
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
		item := models.AuditLog{}
		itemBinded := ctx.BindRequest(item, formData)

		repo := repositories.NewAuditLogRepository(ctx)
		result, err := repo.Create(itemBinded.(models.AuditLog))
		if err != nil {
			ctx.SendError(http.StatusInternalServerError, err)
			return
		}
		ctx.SendResponse(http.StatusOK, result)

	})

	app.Put("/audit-log/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "AuditLog.Put")
			defer ctx.Close()

			repo := repositories.NewAuditLogRepository(ctx)
			id, _ := ctx.GetParamInt64("id")
			item, err := repo.FindOneByID(id)
			if err != nil {
				ctx.SendError(http.StatusNotFound, err)
				return
			}

			formData := forms.AuditLogForm{}
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

	app.Delete("/audit-log/{id:int}",
		func(c iris.Context) {

			ctx := context.NewIris(c, "AuditLog.Delete")
			defer ctx.Close()

			repo := repositories.NewAuditLogRepository(ctx)
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
