package forms

import (
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//AccountForm ...
type AccountForm struct {
	ID          int64      ` form:"id" valid:"-"`
	UserID      int64      ` form:"user_id" valid:"required"`
	Label       string     ` form:"label" valid:"required"`
	Username    string     ` form:"username" valid:"required"`
	Hint        string     ` form:"hint,omitempty" valid:"-"`
	Password    string     ` form:"password" valid:"required"`
	DateCreated time.Time  ` form:"date_created" valid:"required"`
	DateUpdated *time.Time ` form:"date_updated,omitempty" valid:"-"`
}

//IsValid ...
func (a *AccountForm) IsValid() error {
	return validator.Validate(a)
}
