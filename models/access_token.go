package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//AccessToken ...
type AccessToken struct {
	ID          int64      `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:id_UNIQUE;column:id"`
	DateCreated time.Time  `valid:"-" gorm:"type:datetime;not null;column:date_created"`
	DateExpire  *time.Time `valid:"-" gorm:"type:datetime;column:date_expire"`
	Token       string     `valid:"required" gorm:"type:varchar(500);unique;unique_index:token_UNIQUE;index:token_IDX;column:token"`
	UserID      int64      `valid:"required" gorm:"type:bigint(20);not null;index:fk_access_token_1_idx;column:user_id"`

	User User `valid:"-" gorm:"foreignkey:UserID"`
}

//AccessTokenList ...
type AccessTokenList struct {
	Total  int
	Limit  int
	Offset int
	Items  []AccessToken
}

//TableName ...
func (AccessToken) TableName() string {
	return "access_token"
}

//IsValid ...
func (a *AccessToken) IsValid() error {
	return validator.Validate(a)
}

//BeforeCreate ...
func (a *AccessToken) BeforeCreate(scope *gorm.Scope) error {
	created := time.Now()
	scope.SetColumn("date_created", created)

	return a.IsValid()
}

//BeforeUpdate ...
func (a *AccessToken) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}
