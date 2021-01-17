package models

type Comment struct {
	ID        int    `json:"id"`
	PostID    int    `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsAdmin   bool   `json:"is_admin"`
	OPEmail   string `json:"op_email"`
	OPName    string `json:"op_name"`
}
