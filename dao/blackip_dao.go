package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/iralance/go-lottery/models"
	"log"
)

type BlackipDao struct {
	engine *xorm.Engine
}

func NewBlackipDao(engine *xorm.Engine) *BlackipDao {
	return &BlackipDao{
		engine: engine,
	}
}

func (d *BlackipDao) Get(id int) *models.LtBlackip {
	data := &models.LtBlackip{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	return nil
}

func (d *BlackipDao) GetAll(page, size int) []models.LtBlackip {
	offset := (page - 1) * size
	dataList := make([]models.LtBlackip, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&dataList)
	if err != nil {
		log.Println("black_dao.GetALl error=", err)
		return dataList
	}
	return dataList
}

func (d *BlackipDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtBlackip{})
	if err != nil {
		return 0
	}
	return num
}

func (d *BlackipDao) Search(ip string) []models.LtBlackip {
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.
		Where("ip=?", ip).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *BlackipDao) Delete(id int) error {
	data := &models.LtBlackip{Id: id}
	_, err := d.engine.Id(data.Id).
		Update(data)
	return err
}

func (d *BlackipDao) Update(data *models.LtBlackip, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackipDao) Create(data *models.LtBlackip) (int64, error) {
	return d.engine.Insert(data)
}

func (d *BlackipDao) GetByIp(ip string) *models.LtBlackip {
	data := &models.LtBlackip{}
	ok, err := d.engine.Desc("id").Where("ip = ?", ip).Get(data)
	if ok && err == nil {
		return data
	}
	return nil
}
