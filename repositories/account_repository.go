package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
)

//AccountRepository ...
type AccountRepository struct {
	db.BaseRepository
}

//NewAccountRepository ...
func NewAccountRepository(ctx context.Context) AccountRepository {
	return AccountRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//ListByUserID ...
func (a AccountRepository) ListByUserID(userID int64, offset, limit int, sort, sortType string) (list models.AccountList, err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.AccountList{Limit: limit, Offset: offset}
	qb := a.preload(cn.Model(models.Account{})).
		Offset(offset).
		Limit(limit).
		Where("user_id = ?", userID)

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

//List ...
func (a AccountRepository) List(offset, limit int, sort, sortType string) (list models.AccountList, err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.AccountList{Limit: limit, Offset: offset}
	qb := a.preload(cn.Model(models.Account{}))
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
func (a AccountRepository) SearchList(query string, offset, limit int) (list models.AccountList, err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()

	list = models.AccountList{Limit: limit, Offset: offset}
	qb := a.preload(cn.Model(models.Account{})).
		Offset(offset).
		Limit(limit).
		Where("username LIKE ?", "%"+query+"%")

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
func (a AccountRepository) FindOneByID(ID int64) (*models.Account, error) {
	return a.FindOne("id = ?", ID)
}

//FindOneByUsername ...
func (a AccountRepository) FindOneByUsername(username string) (*models.Account, error) {
	return a.FindOne("username = ?", username)
}

//Create ...
func (a AccountRepository) Create(object models.Account) (*models.Account, error) {

	cn, err := a.DB()
	if err != nil {
		return nil, err
	}
	defer a.Close()

	err = cn.Create(&object).Error
	if err != nil {
		return nil, err
	}

	return &object, nil

}

//UpdateColumns ...
func (a AccountRepository) UpdateColumns(object models.Account, values interface{}) (err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()

	err = cn.Model(&object).Updates(values).Error
	if err != nil {
		return
	}

	return
}

//UpdateSingle ...
func (a AccountRepository) UpdateSingle(object models.Account, column string, value interface{}) (err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()

	err = cn.Model(&object).Update(column, value).Error
	if err != nil {
		return
	}

	return
}

//Delete ...
func (a AccountRepository) Delete(object models.Account) (err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()

	err = cn.Where("id = ? ", object.ID).Delete(object).Error
	if err != nil {
		return
	}
	return
}

//FindOne ...
func (a AccountRepository) FindOne(query interface{}, args ...interface{}) (*models.Account, error) {

	cn, err := a.DB()
	if err != nil {
		return nil, err
	}
	defer a.Close()

	var result models.Account
	err = a.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a AccountRepository) preload(query *gorm.DB) *gorm.DB {
	return query.Preload("User")
}
