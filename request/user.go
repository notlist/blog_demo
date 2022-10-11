package request

type UserAddReq struct {
	Name       string `json:"name" `
	NickName   string `json:"nick_name" `
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	DelStatus  int    `json:"del_status"`
	UpdateTime int64  `json:"update_time"`
	CreateTime int64  `json:"create_time"`
}
