package controllers

import (
	"github.com/iralance/go-lottery/models"
	"github.com/iralance/go-lottery/services"
	"time"
)

func (api *LuckyApi) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := services.NewUserService().Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		// 黑名单存在并且有效
		return false, info
	}
	return true, info
}
