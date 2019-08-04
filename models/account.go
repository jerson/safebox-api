package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//Account ...
type Account struct {
	ID          int64      `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:id_UNIQUE;column:id"`
	UserID      int64      `valid:"required" gorm:"type:bigint(20);not null;index:fk_account_1_idx;column:user_id"`
	Label       string     `valid:"runelength(1|50)~Label must have at least 1 character,required~Label is required" gorm:"type:varchar(250);not null;column:label"`
	Username    string     `valid:"required~Username is required" gorm:"type:varchar(250);not null;column:username"`
	Hint        string     `valid:"-" gorm:"type:varchar(250);column:hint"`
	Password    string     `valid:"required~Password is required" gorm:"type:text;not null;column:password"`
	DateCreated time.Time  `valid:"-" gorm:"type:datetime;not null;column:date_created"`
	DateUpdated *time.Time `valid:"-" gorm:"type:datetime;column:date_updated"`

	User User `valid:"-" gorm:"foreignkey:UserID"`
}

//AccountList ...
type AccountList struct {
	Total  int
	Limit  int
	Offset int
	Items  []Account
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
	return nil
}
