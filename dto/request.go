package dto

type UserAddReq struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name" `
	Password string `json:"password"`
	Email    string `json:"email"`
}

type BlogListReq struct {
	Tags []string `json:"tags"`
}
