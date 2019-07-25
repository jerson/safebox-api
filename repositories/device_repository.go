package repositories

import (
	"github.com/jinzhu/gorm"
	"safebox.jerson.dev/api/models"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/modules/db"
	"safebox.jerson.dev/api/modules/util"
)

//DeviceRepository ...
type DeviceRepository struct {
	db.BaseRepository
}

//NewDeviceRepository ...
func NewDeviceRepository(ctx context.Context) DeviceRepository {
	return DeviceRepository{BaseRepository: db.NewBaseRepository(ctx)}
}

//ListByUserID ...
func (d DeviceRepository) ListByUserID(userID int64, offset, limit int, sort, sortType string) (list models.DeviceList, err error) {

	cn, err := d.DB()
	if err != nil {
		return
	}
	defer d.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.DeviceList{Limit: limit, Offset: offset}
	qb := d.preload(cn.Model(models.Device{})).
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
func (d DeviceRepository) List(offset, limit int, sort, sortType string) (list models.DeviceList, err error) {

	cn, err := d.DB()
	if err != nil {
		return
	}
	defer d.Close()
	sortAllows := map[string]string{
		"id": "id",
	}

	list = models.DeviceList{Limit: limit, Offset: offset}
	qb := d.preload(cn.Model(models.Device{}))
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
func (d DeviceRepository) SearchList(query string, offset, limit int) (list models.DeviceList, err error) {

	cn, err := d.DB()
	if err != nil {
		return
	}
	defer d.Close()

	list = models.DeviceList{Limit: limit, Offset: offset}
	qb := d.preload(cn.Model(models.Device{})).
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
func (d DeviceRepository) FindOneByID(ID int64) (*models.Device, error) {
	return d.FindOne("id = ?", ID)
}

//FindOneByHash ...
func (d DeviceRepository) FindOneByHash(hash string) (*models.Device, error) {
	return d.FindOne("hash = ?", hash)
}

//Create ...
func (d DeviceRepository) Create(object models.Device) (*models.Device, error) {

	cn, err := d.DB()
	if err != nil {
		return nil, err
	}
	defer d.Close()

	err = cn.Create(&object).Error
	if err != nil {
		return nil, err
	}

	return &object, nil

}

//UpdateColumns ...
func (d DeviceRepository) UpdateColumns(object models.Device, values interface{}) (err error) {

	cn, err := d.DB()
	if err != nil {
		return
	}
	defer d.Close()

	err = cn.Model(&object).Updates(values).Error
	if err != nil {
		return
	}

	return
}

//UpdateSingle ...
func (d DeviceRepository) UpdateSingle(object models.Device, column string, value interface{}) (err error) {

	cn, err := d.DB()
	if err != nil {
		return
	}
	defer d.Close()

	err = cn.Model(&object).Update(column, value).Error
	if err != nil {
		return
	}

	return
}

//Delete ...
func (d DeviceRepository) Delete(object models.Device) (err error) {

	cn, err := d.DB()
	if err != nil {
		return
	}
	defer d.Close()

	err = cn.Where("id = ? ", object.ID).Delete(object).Error
	if err != nil {
		return
	}
	return
}

//FindOne ...
func (d DeviceRepository) FindOne(query interface{}, args ...interface{}) (*models.Device, error) {

	cn, err := d.DB()
	if err != nil {
		return nil, err
	}
	defer d.Close()

	var result models.Device
	err = d.preload(cn.Where(query, args...)).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (d DeviceRepository) preload(query *gorm.DB) *gorm.DB {
	return query.Preload("User")
}
