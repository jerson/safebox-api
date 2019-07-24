package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/validator"
	"time"
)

//Purchase ...
type Purchase struct {
	ID        int64     `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:idpurchase_UNIQUE;column:id" json:"id"`
	UserID    int64     `valid:"required" gorm:"type:bigint(20);not null;unique_index:user_id_UNIQUE;column:user_id" json:"user_id"`
	ProductID int64     `valid:"required" gorm:"type:bigint(20);not null;unique_index:user_id_UNIQUE;index:fk_purchase_2_idx;column:product_id" json:"product_id"`
	Payload   string    `valid:"required" gorm:"type:text;column:payload" json:"payload,omitempty"`
	Date      time.Time `valid:"-" gorm:"type:datetime;not null;column:date" json:"date"`

	User    User    `valid:"-" gorm:"foreignkey:UserID" json:"user"`
	Product Product `valid:"-" gorm:"foreignkey:ProductID" json:"product"`
}

//PurchaseList ...
type PurchaseList struct {
	Total  int        `json:"total"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Items  []Purchase `json:"items"`
}

//TableName ...
func (Purchase) TableName() string {
	return "purchase"
}

//IsValid ...
func (p *Purchase) IsValid() error {
	return validator.Validate(p)
}

//BeforeCreate ...
func (p *Purchase) BeforeCreate(scope *gorm.Scope) error {

	return p.IsValid()
}

//BeforeUpdate ...
func (p *Purchase) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}
