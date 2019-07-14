package forms

import (
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//UserForm ...
type UserForm struct {
	ID            int64      ` form:"id" valid:"-"`
	Username      string     ` form:"username" valid:"required"`
	PrivateKey    string     ` form:"private_key" valid:"required"`
	DateCreated   time.Time  ` form:"date_created" valid:"required"`
	DateConnected *time.Time ` form:"date_connected,omitempty" valid:"-"`
	PublicKey     string     ` form:"public_key" valid:"required"`
}

//IsValid ...
func (u *UserForm) IsValid() error {
	return validator.Validate(u)
}
