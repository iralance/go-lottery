package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/iralance/go-lottery/models"
	"log"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (d *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	return nil
}

func (d *UserDao) GetAll(page, size int) []models.LtUser {
	offset := (page - 1) * size
	dataList := make([]models.LtUser, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&dataList)
	if err != nil {
		log.Println("user_dao.GetAll error=", err)
		return dataList
	}
	return dataList
}

func (d *UserDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUser{})
	if err != nil {
		return 0
	}
	return num
}

func (d *UserDao) Delete(id int) error {
	data := &models.LtUser{Id: id}
	_, err := d.engine.Id(data.Id).
		Update(data)
	return err
}

func (d *UserDao) Update(data *models.LtUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDao) Create(data *models.LtUser) (int64, error) {
	return d.engine.Insert(data)
}
