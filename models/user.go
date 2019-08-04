package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//User ...
type User struct {
	ID              int64      `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;column:id"`
	Username        string     `valid:"runelength(4|30)~Username must be at least 4 characters,required~Username is required" gorm:"type:varchar(45);unique;not null;unique_index:username_UNIQUE;column:username"`
	PrivateKey      string     `valid:"required~Error generating keys for password" gorm:"type:text;not null;column:private_key"`
	PublicKey       string     `valid:"required~Error generating keys for password" gorm:"type:text;not null;column:public_key"`
	Email           string     `valid:"email~Invalid email,optional" gorm:"type:varchar(191);column:email"`
	LocationEnabled bool       `valid:"optional" gorm:"type:tinyint(1);column:location_enabled"`
	DateCreated     time.Time  `valid:"-" gorm:"type:datetime;not null;column:date_created"`
	DateConnected   *time.Time `valid:"-" gorm:"type:datetime;column:date_connected"`
}

//UserList ...
type UserList struct {
	Total  int
	Limit  int
	Offset int
	Items  []User
}

//TableName ...
func (User) TableName() string {
	return "user"
}

//IsValid ...
func (u *User) IsValid() error {
	return validator.Validate(u)
}

//BeforeCreate ...
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	created := time.Now()
	scope.SetColumn("date_created", created)

	return u.IsValid()
}

//BeforeUpdate ...
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}
