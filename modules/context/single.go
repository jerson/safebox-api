package context

import (
	"context"
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/util"
)

// AppContext ...
type AppContext struct {
	template
	Context context.Context
	Data    map[string]interface{}
}

// NewContext ...
func NewContext(context context.Context, category string) *AppContext {
	ctx := &AppContext{Context: context, template: template{Category: category}}
	ctx.init()
	return ctx
}

// NewContextSingle ...
func NewContextSingle(category string) *AppContext {
	ctx := &AppContext{Context: context.Background(), template: template{Category: category}}
	ctx.init()
	return ctx
}

// init ...
func (r *AppContext) init() {
	r.Data = map[string]interface{}{}
	r.Token = util.UniqueID()
}

// Close ...
func (r *AppContext) Close() {
	db := r.Get("DB")
	if db != nil {
		cn := db.(*gorm.DB)
		cn.Close()
		if config.Vars.Debug {
			r.GetLogger("DB").Debug("conexion cerrada")
		}
	}
}

// Set ...
func (r *AppContext) Set(name string, value interface{}) {
	r.Data[name] = value
}

// Get ...
func (r AppContext) Get(name string) interface{} {
	return r.Data[name]
}
