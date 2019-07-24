package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
)

//AuditLog ...
type AuditLog struct {
	ID      int64  `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:idaudit_log_UNIQUE;column:id" json:"id"`
	UserID  *int64 `valid:"-" gorm:"type:bigint(20);index:fk_audit_log_1_idx;column:user_id" json:"user_id,omitempty"`
	IP      string `valid:"-" gorm:"type:varchar(45);not null;index:idx_ip;column:ip" json:"ip"`
	Action  string `valid:"required" gorm:"type:varchar(45);not null;index:idx_action;column:action" json:"action"`
	Payload string `valid:"required" gorm:"type:text;column:payload" json:"payload,omitempty"`

	User *User `gorm:"foreignkey:UserID" json:"user,omitempty"`
}

//AuditLogList ...
type AuditLogList struct {
	Total  int        `json:"total"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Items  []AuditLog `json:"items"`
}

//TableName ...
func (AuditLog) TableName() string {
	return "audit_log"
}

//IsValid ...
func (a *AuditLog) IsValid() error {
	return validator.Validate(a)
}

//BeforeCreate ...
func (a *AuditLog) BeforeCreate(scope *gorm.Scope) error {

	return a.IsValid()
}

//BeforeUpdate ...
func (a *AuditLog) BeforeUpdate(scope *gorm.Scope) error {
	return a.IsValid()
}
