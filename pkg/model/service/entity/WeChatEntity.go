package entity

type UserInfo struct {
	OpenId     string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Unionid    string `json:"unionid,omitempty"`
	Errcode    int    `json:"errcode,omitempty"`
	Errmsg     string `json:"errmsg,omitempty"`
} 




