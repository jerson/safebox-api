package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
)

//PurchaseRepository ...
type PurchaseRepository struct {
	db.BaseRepository
}

//NewPurchaseRepository ...
func NewPurchaseRepository(ctx context.Context) PurchaseRepository {
	return PurchaseRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//List ...
func (p PurchaseRepository) List(offset, limit int, sort, sortType string) (list models.PurchaseList, err error) {

	cn, err := p.DB()
	if err != nil {
		return
	}
	defer p.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.PurchaseList{Limit: limit, Offset: offset}
	qb := p.preload(cn.Model(models.Purchase{}))
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

//FindOneByID ...
func (p PurchaseRepository) FindOneByID(ID int64) (*models.Purchase, error) {
	return p.FindOne("id = ?", ID)
}

//FindOneByUser ...
func (p PurchaseRepository) FindOneByUser(userID int64, productID int64) (*models.Purchase, error) {
	return p.FindOne("user_id = ? AND product_id", userID, productID)
}

//Create ...
func (p PurchaseRepository) Create(object models.Purchase) (*models.Purchase, error) {

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

//UpdateColumns ...
func (p PurchaseRepository) UpdateColumns(object models.Purchase, values interface{}) (err error) {

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
func (p PurchaseRepository) UpdateSingle(object models.Purchase, column string, value interface{}) (err error) {

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
func (p PurchaseRepository) Delete(object models.Purchase) (err error) {

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
func (p PurchaseRepository) FindOne(query interface{}, args ...interface{}) (*models.Purchase, error) {

	cn, err := p.DB()
	if err != nil {
		return nil, err
	}
	defer p.Close()

	var result models.Purchase
	err = p.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p PurchaseRepository) preload(query *gorm.DB) *gorm.DB {
	return query.Preload("User").Preload("Product")
}
