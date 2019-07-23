package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//AccessToken ...
type AccessToken struct {
	ID          int64      `gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:id_UNIQUE;column:id" json:"id"`
	DateCreated time.Time  `gorm:"type:datetime;not null;column:date_created" json:"date_created"`
	DateExpire  *time.Time `gorm:"type:datetime;column:date_expire" json:"date_expire,omitempty"`
	Token       string     `gorm:"type:varchar(500);unique;unique_index:token_UNIQUE;index:token_IDX;column:token" json:"token,omitempty"`
	UserID      int64      `gorm:"type:bigint(20);not null;index:fk_access_token_1_idx;column:user_id" json:"user_id"`

	User User `gorm:"foreignkey:UserID" json:"user"`
}

//AccessTokenList ...
type AccessTokenList struct {
	Total  int           `json:"total"`
	Limit  int           `json:"limit"`
	Offset int           `json:"offset"`
	Items  []AccessToken `json:"items"`
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
	return a.IsValid()
}
