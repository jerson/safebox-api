package forms

import (
	"safebox.jerson.dev/api/modules/validator"
)

//ProductForm ...
type ProductForm struct {
	ID          int64  ` form:"id" valid:"-"`
	Slug        string ` form:"slug" valid:"required"`
	Name        string ` form:"name" valid:"required"`
	Description string ` form:"description,omitempty" valid:"-"`
}

//IsValid ...
func (p *ProductForm) IsValid() error {
	return validator.Validate(p)
}
