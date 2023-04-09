package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func GetBusScore(busId string) (sumScore float32, sumBaby int64, err error) {
	return dao.GetBusScore(busId)
}

func SelectStations(keyWords string) ([]model.Station, error) {
	return dao.LikeSelect(keyWords)
}

func GetStationDetails(stationName string) (*model.Station, error) {
	return dao.GetStationDetails(stationName)
}

func StationsScoreIncr(IncrNum float64, stationName string) error {
	return dao.StationsScoreIncr(IncrNum, stationName)
}

func GetHot() []string {
	return dao.GetHotStations()
}

//防止恶意刷新导致热榜失真

// CheckLimit 查看此时此ip是否超出访问次数
func CheckLimit(ip string) error {
	count, err := dao.CheckLimit(ip)
	if err != nil {
		return err
	}
	if count > config.LimitedRequest {
		return config.TooManyRequests
	}
	return nil
}

func IpRefresh(ip string) (err error) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	if err = dao.IpIncrBy(ip); err != nil {
		return err
	}

	//设置过期时间
	return dao.TimeExpire(ip)
}
