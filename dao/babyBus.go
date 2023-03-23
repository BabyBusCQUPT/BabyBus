package dao

import "BabyBus/model"

func SaveScore(babyBus *model.BabyBus) error {
	if err := DB.Model(&model.BabyBus{}).Create(&babyBus).Error; err != nil {
		return err
	}
	return nil
}

func GetScore(bus *model.BabyBus) (scores []model.BabyBus, err error) {
	if err = DB.Model(model.BabyBus{}).Where("babyId = ? AND busID = ?", bus.BabyId, bus.BabyId).Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}
