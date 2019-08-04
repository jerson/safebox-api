package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
)

//UserLocationRepository ...
type UserLocationRepository struct {
	db.BaseRepository
}

//NewUserLocationRepository ...
func NewUserLocationRepository(ctx context.Context) UserLocationRepository {
	return UserLocationRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//List ...
func (u UserLocationRepository) List(offset, limit int, sort, sortType string) (list models.UserLocationList, err error) {

	cn, err := u.DB()
	if err != nil {
		return
	}
	defer u.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.UserLocationList{Limit: limit, Offset: offset}
	qb := u.preload(cn.Model(models.UserLocation{}))
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
func (u UserLocationRepository) FindOneByID(ID int64) (*models.UserLocation, error) {
	return u.FindOne("id = ?", ID)
}

//Create ...
func (u UserLocationRepository) Create(object models.UserLocation) (*models.UserLocation, error) {

	cn, err := u.DB()
	if err != nil {
		return nil, err
	}
	defer u.Close()

	err = cn.Create(&object).Error
	if err != nil {
		return nil, err
	}

	return &object, nil

}

//UpdateColumns ...
func (u UserLocationRepository) UpdateColumns(object models.UserLocation, values interface{}) (err error) {

	cn, err := u.DB()
	if err != nil {
		return
	}
	defer u.Close()

	err = cn.Model(&object).Updates(values).Error
	if err != nil {
		return
	}

	return
}

//UpdateSingle ...
func (u UserLocationRepository) UpdateSingle(object models.UserLocation, column string, value interface{}) (err error) {

	cn, err := u.DB()
	if err != nil {
		return
	}
	defer u.Close()

	err = cn.Model(&object).Update(column, value).Error
	if err != nil {
		return
	}

	return
}

//Delete ...
func (u UserLocationRepository) Delete(object models.UserLocation) (err error) {

	cn, err := u.DB()
	if err != nil {
		return
	}
	defer u.Close()

	err = cn.Where("id = ? ", object.ID).Delete(object).Error
	if err != nil {
		return
	}
	return
}

//FindOne ...
func (u UserLocationRepository) FindOne(query interface{}, args ...interface{}) (*models.UserLocation, error) {

	cn, err := u.DB()
	if err != nil {
		return nil, err
	}
	defer u.Close()

	var result models.UserLocation
	err = u.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u UserLocationRepository) preload(query *gorm.DB) *gorm.DB {
	return query.Preload("User")
}
