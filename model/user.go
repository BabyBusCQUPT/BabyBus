package model

type User struct {
	//gorm.Model软删除
	Id            string
	OpenId        string
	SessionKey    string
	Age           int
	Gender        byte
	Token         string
	NikeName      string
	Image         string
	UserLongitude float32
	UserLatitude  float32
	MostUsed      int
}
