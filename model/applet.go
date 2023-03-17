package model

type Applet struct {
	AppId     string
	AppSecret string
	GrantType string
	BasicUrl  string
}

type WeChatConnection struct {
	Id         int    //token
	OpenId     string //token
	SessionKey string //token
	ErrMsg     string
	ErrCode    int32
}
