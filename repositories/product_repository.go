package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
)

//ProductRepository ...
type ProductRepository struct {
	db.BaseRepository
}

//NewProductRepository ...
func NewProductRepository(ctx context.Base) ProductRepository {
	return ProductRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//List ...
func (p ProductRepository) List(offset, limit int, sort, sortType string) (list models.ProductList, err error) {

	cn, err := p.DB()
	if err != nil {
		return
	}
	defer p.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.ProductList{Limit: limit, Offset: offset}
	qb := p.preload(cn.Model(models.Product{}))
	err = qb.
		Order(util.SortValues(sort, sortType, sortAllows)).
		Offset(offset).
		Limit(limit).
		Find(&list.Items).
		Error

	if err != nil {
		return
	}
	err = qb.Count(&list.Total).Error
	if err != nil {
		list.Total = len(list.Items)
	}

	return
}

//SearchList ...
func (p ProductRepository) SearchList(query string, offset, limit int) (list models.ProductList, err error) {

	cn, err := p.DB()
	if err != nil {
		return
	}
	defer p.Close()

	list = models.ProductList{Limit: limit, Offset: offset}
	qb := p.preload(cn.Model(models.Product{})).
		Offset(offset).
		Limit(limit).
		Where("name LIKE ?", "%"+query+"%")

	err = qb.Find(&list.Items).
		Error

	if err != nil {
		return
	}
	err = qb.Count(&list.Total).Error
	if err != nil {
		list.Total = len(list.Items)
	}

	return
}

//FindOneByID ...
func (p ProductRepository) FindOneByID(ID int64) (*models.Product, error) {
	return p.FindOne("id = ?", ID)
}

//FindOneByName ...
func (p ProductRepository) FindOneByName(name string) (*models.Product, error) {
	return p.FindOne("name = ?", name)
}

//Create ...
func (p ProductRepository) Create(object models.Product) (*models.Product, error) {

	cn, err := p.DB()
	if err != nil {
		return nil, err
	}
	defer p.Close()

	err = cn.Create(&object).Error
	if err != nil {
		return nil, err
	}

	return &object, nil

}

//Update ...
func (p ProductRepository) Update(object models.Product, data interface{}) (err error) {
	values := p.DiffStruct(&object, data)
	return p.UpdateColumns(object, values)
}

//UpdateColumns ...
func (p ProductRepository) UpdateColumns(object models.Product, values interface{}) (err error) {

	cn, err := p.DB()
	if err != nil {
		return
	}
	defer p.Close()

	err = cn.Model(&object).Updates(values).Error
	if err != nil {
		return
	}

	return
}

//UpdateSingle ...
func (p ProductRepository) UpdateSingle(object models.Product, column string, value interface{}) (err error) {

	cn, err := p.DB()
	if err != nil {
		return
	}
	defer p.Close()

	err = cn.Model(&object).Update(column, value).Error
	if err != nil {
		return
	}

	return
}

//Delete ...
func (p ProductRepository) Delete(object models.Product) (err error) {

	cn, err := p.DB()
	if err != nil {
		return
	}
	defer p.Close()

	err = cn.Where("id = ? ", object.ID).Delete(object).Error
	if err != nil {
		return
	}
	return
}

//FindOne ...
func (p ProductRepository) FindOne(query interface{}, args ...interface{}) (*models.Product, error) {

	cn, err := p.DB()
	if err != nil {
		return nil, err
	}
	defer p.Close()

	var result models.Product
	err = p.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p ProductRepository) preload(query *gorm.DB) *gorm.DB {
	return query
}
