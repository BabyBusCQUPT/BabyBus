package service

import "BabyBus/dao"

func GetBusScore(busId string) (sumScore float32, sumBaby int64, err error) {
	return dao.GetBusScore(busId)
}
