package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
	"time"
)

//AccessTokenRepository ...
type AccessTokenRepository struct {
	db.BaseRepository
}

//NewAccessTokenRepository ...
func NewAccessTokenRepository(ctx context.Context) AccessTokenRepository {
	return AccessTokenRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//List ...
func (a AccessTokenRepository) List(offset, limit int, sort, sortType string) (list models.AccessTokenList, err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.AccessTokenList{Limit: limit, Offset: offset}
	qb := a.preload(cn.Model(models.AccessToken{}))
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
func (a AccessTokenRepository) FindOneByID(ID int64) (*models.AccessToken, error) {
	return a.FindOne("id = ?", ID)
}

//FindOneByToken ...
func (a AccessTokenRepository) FindOneByToken(token string) (*models.AccessToken, error) {
	return a.FindOne("token = ?", token)
}

//FindOnyByUserID ...
func (a AccessTokenRepository) FindOnyByUserID(userID int64) (*models.AccessToken, error) {
	return a.FindOne("user_id = ?", userID)
}

//Create ...
func (a AccessTokenRepository) Create(object models.AccessToken) (*models.AccessToken, error) {

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
func (a AccessTokenRepository) UpdateColumns(object models.AccessToken, values interface{}) (err error) {

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
func (a AccessTokenRepository) UpdateSingle(object models.AccessToken, column string, value interface{}) (err error) {

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
func (a AccessTokenRepository) Delete(object models.AccessToken) (err error) {

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

//DeleteExpired ...
func (a AccessTokenRepository) DeleteExpired() (err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()

	err = cn.Where("date_expire < ? ", time.Now()).Delete(models.AccessToken{}).Error
	if err != nil {
		return
	}
	return
}

//FindOne ...
func (a AccessTokenRepository) FindOne(query interface{}, args ...interface{}) (*models.AccessToken, error) {

	cn, err := a.DB()
	if err != nil {
		return nil, err
	}
	defer a.Close()

	var result models.AccessToken
	err = a.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a AccessTokenRepository) preload(query *gorm.DB) *gorm.DB {
	return query.Preload("User")
}
