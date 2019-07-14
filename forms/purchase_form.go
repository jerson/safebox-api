package forms

import (
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//PurchaseForm ...
type PurchaseForm struct {
	ID        int64     ` form:"id" valid:"-"`
	UserID    int64     ` form:"user_id" valid:"required"`
	ProductID int64     ` form:"product_id" valid:"required"`
	Payload   string    ` form:"payload,omitempty" valid:"-"`
	Date      time.Time ` form:"date" valid:"required"`
}

//IsValid ...
func (p *PurchaseForm) IsValid() error {
	return validator.Validate(p)
}
