package entity

type EmojiFile struct {
	Id            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Extension     string    `json:"extension,omitempty"`
	Path          string    `json:"path,omitempty"`
	Status        int       `json:"status,omitempty"`
	CreateTime    string    `json:"create_time,omitempty"`
	UpdateTime    string    `json:"update_time,omitempty"`
	Md5Encode     string    `json:"md5_encode,omitempty"`
	ImageUrl      string    `json:"image_url,omitempty"`
	BasePath      string    `json:"base_path,omitempty"`
	SentenceCount int       `json:"sentence_count,omitempty"`
	Sentence      string    `json:"sentence,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
}

