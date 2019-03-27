package entity

type UserEmojiFile struct {
	Id         int    `json:"id"`
	OpenId     string `json:"open_id"`
	Status     int    `json:"status"`
	ImageUrl   string  `json:"image_url"`
	CreateTime string `json:"create_time"`
}
