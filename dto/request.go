package dto

type UserAddReq struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name" `
	Password string `json:"password"`
	Email    string `json:"email"`
}

type BlogListReq struct {
	UserId int64    `json:"user_id"`
	Tags   []string `json:"tags"`
}

type BlogCreateReq struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type BlogEditReq struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}
