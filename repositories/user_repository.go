package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
)

//UserRepository ...
type UserRepository struct {
	db.BaseRepository
}

//NewUserRepository ...
func NewUserRepository(ctx context.Context) UserRepository {
	return UserRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//List ...
func (u UserRepository) List(offset, limit int, sort, sortType string) (list models.UserList, err error) {

	cn, err := u.DB()
	if err != nil {
		return
	}
	defer u.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.UserList{Limit: limit, Offset: offset}
	qb := u.preload(cn.Model(models.User{}))
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
func (u UserRepository) SearchList(query string, offset, limit int) (list models.UserList, err error) {

	cn, err := u.DB()
	if err != nil {
		return
	}
	defer u.Close()

	list = models.UserList{Limit: limit, Offset: offset}
	qb := u.preload(cn.Model(models.User{})).
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
func (u UserRepository) FindOneByID(ID int64) (*models.User, error) {
	return u.FindOne("id = ?", ID)
}

//FindOneByUsername ...
func (u UserRepository) FindOneByUsername(username string) (*models.User, error) {
	return u.FindOne("username = ?", username)
}

//Create ...
func (u UserRepository) Create(object models.User) (*models.User, error) {

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
func (u UserRepository) UpdateColumns(object models.User, values interface{}) (err error) {

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
func (u UserRepository) UpdateSingle(object models.User, column string, value interface{}) (err error) {

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
func (u UserRepository) Delete(object models.User) (err error) {

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
func (u UserRepository) FindOne(query interface{}, args ...interface{}) (*models.User, error) {

	cn, err := u.DB()
	if err != nil {
		return nil, err
	}
	defer u.Close()

	var result models.User
	err = u.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u UserRepository) preload(query *gorm.DB) *gorm.DB {
	return query
}
