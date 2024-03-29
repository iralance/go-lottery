package services

import (
	"github.com/iralance/go-lottery/dao"
	"github.com/iralance/go-lottery/datasource"
	"github.com/iralance/go-lottery/models"
)

type ResultService interface {
	GetAll(page, size int) []models.LtResult
	CountAll() int64
	GetNewPrize(size int, giftIds []int) []models.LtResult
	SearchByGift(giftId, page, size int) []models.LtResult
	SearchByUser(uid, page, size int) []models.LtResult
	CountByGift(giftId int) int64
	CountByUser(uid int) int64
	Get(id int) *models.LtResult
	Delete(id int) error
	Update(data *models.LtResult, columns []string) error
	Create(data *models.LtResult) (int64, error)
}

type resultService struct {
	dao *dao.ResultDao
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(datasource.InstanceDbMaster()),
	}
}

func (g *resultService) GetAll(page, size int) []models.LtResult {
	return g.dao.GetAll(page, size)
}

func (g *resultService) CountAll() int64 {
	return g.dao.CountAll()
}

func (s *resultService) GetNewPrize(size int, giftIds []int) []models.LtResult {
	return s.dao.GetNewPrize(size, giftIds)
}

func (s *resultService) SearchByGift(giftId, page, size int) []models.LtResult {
	return s.dao.SearchByGift(giftId, page, size)
}

func (s *resultService) SearchByUser(uid, page, size int) []models.LtResult {
	return s.dao.SearchByUser(uid, page, size)
}

func (s *resultService) CountByGift(giftId int) int64 {
	return s.dao.CountByGift(giftId)
}

func (s *resultService) CountByUser(uid int) int64 {
	return s.dao.CountByUser(uid)
}

func (g *resultService) Get(id int) *models.LtResult {
	return g.dao.Get(id)
}

func (g *resultService) Delete(id int) error {
	return g.dao.Delete(id)

}

func (g *resultService) Update(data *models.LtResult, columns []string) error {
	return g.dao.Update(data, columns)

}

func (g *resultService) Create(data *models.LtResult) (int64, error) {
	return g.dao.Create(data)
}
