package utils

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strconv"
)

type AuthUserInfo struct {
	Openid   string
	Nickname string
	Gender   int
	Avatar   string
}

type wxAuthVerifyResult struct {
	ErrCode int `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type wxUserInfo struct {
	UnionId  string `json:"unionid"`
	Nickname string `json:"nickname"`
	Gender   int `json:"sex"`
	Avatar   string `json:"headimgurl"`
}

type wbAuthVerifyResult struct {
	Uid    int64 `json:"uid"`
	AppKey string `json:"appkey"`
}

type wbUserInfo struct {
	Id       int64 `json:"id"`
	Nickname string `json:"screen_name"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar_large"`
}

type qqUserInfo struct {
	Ret      int `json:"ret"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Avatar   string `json:"figureurl_qq_1"`
}

func AuthWithWeiXin(wx_openid, wx_token string) (*AuthUserInfo, int) {
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s&lang=zh_CN",
		wx_token, wx_openid))
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL
	}
	var verifyResult wxAuthVerifyResult
	err = json.Unmarshal(body, &verifyResult)
	if err != nil || verifyResult.ErrCode != 0 {
		return nil, ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL
	}

	resp, err = http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		wx_token, wx_openid))
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL
	}
	var userInfo wxUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL
	}
	return &AuthUserInfo{Openid: userInfo.UnionId, Nickname: userInfo.Nickname, Gender: userInfo.Gender, Avatar: userInfo.Avatar}, 0
}

func AuthWithWeiBo(wb_token string) (*AuthUserInfo, int) {
	resp, err := http.Post("https://api.weibo.com/oauth2/get_token_info",
		"application/x-www-form-urlencoded", bytes.NewBuffer([]byte("access_token=" + wb_token)))
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIBO_AUTH_FAIL
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIBO_AUTH_FAIL
	}
	var verifyResult wbAuthVerifyResult
	err = json.Unmarshal(body, &verifyResult)
	if err != nil || verifyResult.Uid <= 0 {
		return nil, ERROR_CODE_AUTH_WEIBO_AUTH_FAIL
	}

	resp, err = http.Get(fmt.Sprintf("https://api.weibo.com/2/users/show.json?access_token=%s&uid=%d",
		wb_token, verifyResult.Uid))
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIBO_AUTH_FAIL
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIBO_AUTH_FAIL
	}
	var userInfo wbUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, ERROR_CODE_AUTH_WEIBO_AUTH_FAIL
	}
	gender := GENDER_UNKNOWN
	if userInfo.Gender == "m" {
		gender = GENDER_MALE
	} else if userInfo.Gender == "f" {
		gender = GENDER_FEMALE
	}
	return &AuthUserInfo{Openid: strconv.FormatInt(userInfo.Id, 10), Nickname: userInfo.Nickname, Gender: gender, Avatar: userInfo.Avatar}, 0
}

func AuthWithQQ(qq_openid, qq_token, qq_oauth_consumer_key string) (*AuthUserInfo, int) {
	resp, err := http.Get(fmt.Sprintf("https://graph.qq.com/user/get_user_info?oauth_consumer_key=%s&access_token=%s&openid=%s&format=json",
		qq_oauth_consumer_key, qq_token, qq_openid))
	if err != nil {
		return nil, ERROR_CODE_AUTH_QQ_AUTH_FAIL
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ERROR_CODE_AUTH_QQ_AUTH_FAIL
	}
	var userInfo qqUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil || userInfo.Ret != 0 {
		return nil, ERROR_CODE_AUTH_QQ_AUTH_FAIL
	}
	gender := GENDER_UNKNOWN
	if userInfo.Gender == "男" {
		gender = GENDER_MALE
	} else if userInfo.Gender == "女" {
		gender = GENDER_FEMALE
	}
	return &AuthUserInfo{Openid: qq_openid, Nickname: userInfo.Nickname, Gender: gender, Avatar: userInfo.Avatar}, 0
}
