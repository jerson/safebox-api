package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//Account ...
type Account struct {
	ID          int64      `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:id_UNIQUE;column:id" json:"id"`
	UserID      int64      `valid:"required" gorm:"type:bigint(20);not null;index:fk_account_1_idx;column:user_id" json:"user_id"`
	Label       string     `valid:"required" gorm:"type:varchar(250);not null;column:label" json:"label"`
	Username    string     `valid:"required" gorm:"type:varchar(250);not null;column:username" json:"username"`
	Hint        string     `valid:"-" gorm:"type:varchar(250);column:hint" json:"hint,omitempty"`
	Password    string     `valid:"required" gorm:"type:text;not null;column:password" json:"password"`
	DateCreated time.Time  `valid:"-" gorm:"type:datetime;not null;column:date_created" json:"date_created"`
	DateUpdated *time.Time `valid:"-" gorm:"type:datetime;column:date_updated" json:"date_updated,omitempty"`

	User User `gorm:"foreignkey:UserID" json:"user"`
}

//AccountList ...
type AccountList struct {
	Total  int       `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Items  []Account `json:"items"`
}

//TableName ...
func (Account) TableName() string {
	return "account"
}

//IsValid ...
func (a *Account) IsValid() error {
	return validator.Validate(a)
}

//BeforeCreate ...
func (a *Account) BeforeCreate(scope *gorm.Scope) error {
	created := time.Now()
	scope.SetColumn("date_created", created)

	return a.IsValid()
}

//BeforeUpdate ...
func (a *Account) BeforeUpdate(scope *gorm.Scope) error {
	updated := time.Now()
	scope.SetColumn("date_updated", &updated)
	return a.IsValid()
}
