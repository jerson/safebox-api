package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//UserLocation ...
type UserLocation struct {
	ID     int64     `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:id_UNIQUE;column:id" `
	UserID int64     `valid:"required" gorm:"type:bigint(20);unique;not null;unique_index:user_id_UNIQUE;column:user_id" `
	Date   time.Time `valid:"-" gorm:"type:datetime;not null;column:date"`

	User User `valid:"-" gorm:"foreignkey:UserID"`
}

//UserLocationList ...
type UserLocationList struct {
	Total  int            `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Items  []UserLocation `json:"items"`
}

//TableName ...
func (UserLocation) TableName() string {
	return "user_location"
}

//IsValid ...
func (u *UserLocation) IsValid() error {
	return validator.Validate(u)
}

//BeforeCreate ...
func (u *UserLocation) BeforeCreate(scope *gorm.Scope) error {
	current := time.Now()
	scope.SetColumn("date", current)
	return u.IsValid()
}

//BeforeUpdate ...
func (u *UserLocation) BeforeUpdate(scope *gorm.Scope) error {
	current := time.Now()
	scope.SetColumn("date", current)
	return u.IsValid()
}
