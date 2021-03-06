package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//Device ...
type Device struct {
	ID          int64     `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:id_UNIQUE;column:id"`
	Name        string    `valid:"required" gorm:"type:varchar(250);not null;column:name"`
	UID         string    `valid:"required" gorm:"type:varchar(250);not null;column:uid"`
	PublicKey   string    `valid:"required" gorm:"type:text;column:public_key"`
	Hash        string    `valid:"-" gorm:"type:varchar(250);unique;not null;unique_index:hash_UNIQUE;index:hash_idx;column:hash"`
	UserID      int64     `valid:"required" gorm:"type:bigint(20);not null;index:fk_device_1_idx;column:user_id"`
	DateCreated time.Time `valid:"-" gorm:"type:datetime;not null;column:date_created"`

	User User `valid:"-" gorm:"foreignkey:UserID"`
}

//DeviceList ...
type DeviceList struct {
	Total  int
	Limit  int
	Offset int
	Items  []Device
}

//TableName ...
func (Device) TableName() string {
	return "device"
}

//IsValid ...
func (d *Device) IsValid() error {
	return validator.Validate(d)
}

//BeforeCreate ...
func (d *Device) BeforeCreate(scope *gorm.Scope) error {
	created := time.Now()
	scope.SetColumn("date_created", created)

	return d.IsValid()
}

//BeforeUpdate ...
func (d *Device) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}
