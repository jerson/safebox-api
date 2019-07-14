package forms

import (
	"safebox.jerson.dev/api/modules/validator"
)

//AuditLogForm ...
type AuditLogForm struct {
	ID      int64  ` form:"id" valid:"-"`
	UserID  *int64 ` form:"user_id,omitempty" valid:"-"`
	IP      string ` form:"ip" valid:"required"`
	Action  string ` form:"action" valid:"required"`
	Payload string ` form:"payload,omitempty" valid:"-"`
}

//IsValid ...
func (a *AuditLogForm) IsValid() error {
	return validator.Validate(a)
}
