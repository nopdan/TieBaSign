package sign

type Like struct {
	Data  Data   `json:"data"`
	Error string `json:"error"`
	No    int    `json:"no"`
}

type Data struct {
	ItbTbs    string      `json:"itb_tbs"`
	LikeForum []LikeForum `json:"like_forum"`
	Tbs       string      `json:"tbs"`
	UID       float64     `json:"uid"`
}

type LikeForum struct {
	FavoType  int     `json:"favo_type"`
	ForumID   float64 `json:"forum_id"`
	ForumName string  `json:"forum_name"`
	IsLike    bool    `json:"is_like"`
	IsSign    int     `json:"is_sign"`
	UserExp   string  `json:"user_exp"`
	UserLevel string  `json:"user_level"`
}
