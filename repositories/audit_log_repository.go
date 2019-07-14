package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
)

//AuditLogRepository ...
type AuditLogRepository struct {
	db.BaseRepository
}

//NewAuditLogRepository ...
func NewAuditLogRepository(ctx context.Base) AuditLogRepository {
	return AuditLogRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//List ...
func (a AuditLogRepository) List(offset, limit int, sort, sortType string) (list models.AuditLogList, err error) {

	cn, err := a.DB()
	if err != nil {
		return
	}
	defer a.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.AuditLogList{Limit: limit, Offset: offset}
	qb := a.preload(cn.Model(models.AuditLog{}))
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
func (a AuditLogRepository) FindOneByID(ID int64) (*models.AuditLog, error) {
	return a.FindOne("id = ?", ID)
}

//Create ...
func (a AuditLogRepository) Create(object models.AuditLog) (*models.AuditLog, error) {

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

//Update ...
func (a AuditLogRepository) Update(object models.AuditLog, data interface{}) (err error) {
	values := a.DiffStruct(&object, data)
	return a.UpdateColumns(object, values)
}

//UpdateColumns ...
func (a AuditLogRepository) UpdateColumns(object models.AuditLog, values interface{}) (err error) {

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
func (a AuditLogRepository) UpdateSingle(object models.AuditLog, column string, value interface{}) (err error) {

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
func (a AuditLogRepository) Delete(object models.AuditLog) (err error) {

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
func (a AuditLogRepository) FindOne(query interface{}, args ...interface{}) (*models.AuditLog, error) {

	cn, err := a.DB()
	if err != nil {
		return nil, err
	}
	defer a.Close()

	var result models.AuditLog
	err = a.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a AuditLogRepository) preload(query *gorm.DB) *gorm.DB {
	return query.Preload("User")
}
