package dao

import "BabyBus/model"

func GetBusScore(busId string) (sumScore float32, sumBaby int64, err error) {
	if err = DB.Model(&model.BabyBus{}).Where("busId = ?", busId).Pluck("SUM(sumScore) as sumScore", &sumScore).Error; err != nil {
		return -1, -1, err
	}
	if err = DB.Model(&model.BabyBus{}).Where("busId = ?", busId).Count(&sumBaby).Error; err != nil {
		return -1, -1, err
	}
	return sumScore, sumBaby, nil
}

func LikeSelect(keyWords string) (stations []model.Station, err error) {
	if err = DB.Model(&model.Station{}).Where("name LIKE " + "%" + keyWords).Find(&stations).Error; err != nil {
		return nil, err
	}
	return stations, nil
}
