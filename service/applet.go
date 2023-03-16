package service

import (
	"BabyBus/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
)

func ParseAppletConfig() (*model.Applet, error) {
	applet := &model.Applet{}
	jsonFile, err := os.Open("config/applet.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonByte, applet)
	return applet, err
}

func ConnectWeChatApi(applet *model.Applet, code string) (*model.WeChatConnection, error) {
	//eg.   https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	WeChatConnection := &model.WeChatConnection{}
	data := url2.Values{}
	data.Set("appid", applet.AppId)
	data.Set("secret", applet.AppSecret)
	data.Set("js_code", code)
	data.Set("grant_type", applet.GrantType)
	url, err := url2.ParseRequestURI(applet.BasicUrl)
	if err != nil {
		return nil, err
	}
	url.RawQuery = data.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, WeChatConnection)
	return WeChatConnection, err
}
