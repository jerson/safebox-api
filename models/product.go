package models

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/modules/util"
	"safebox.jerson.dev/api/modules/validator"
)

//Product ...
type Product struct {
	ID          int64  `valid:"-" gorm:"primary_key;auto_increment;type:bigint(20);not null;unique_index:id_UNIQUE;column:id" json:"id"`
	Slug        string `valid:"-" gorm:"type:varchar(250);unique;not null;unique_index:slug_UNIQUE;column:slug" json:"slug"`
	Name        string `valid:"required" gorm:"type:varchar(45);not null;column:name" json:"name"`
	Description string `valid:"-" gorm:"type:text;column:description" json:"description,omitempty"`
}

//ProductList ...
type ProductList struct {
	Total  int       `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Items  []Product `json:"items"`
}

//TableName ...
func (Product) TableName() string {
	return "product"
}

//IsValid ...
func (p *Product) IsValid() error {
	return validator.Validate(p)
}

//BeforeCreate ...
func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	if p.Slug == "" {
		scope.SetColumn("slug", util.Slug(p.Name))
	}

	return p.IsValid()
}

//BeforeUpdate ...
func (p *Product) BeforeUpdate(scope *gorm.Scope) error {
	if p.Slug == "" {
		scope.SetColumn("slug", util.Slug(p.Name))
	}
	return nil
}
