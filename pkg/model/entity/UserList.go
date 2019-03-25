package entity

type UserList struct {
	Id        int    `json:"id"`
	OpenId    string `json:"open_id"`
	NickName  string `json:"nick_name"`
	Avatar    string `json:",omitempty"`
}