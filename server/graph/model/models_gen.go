// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthPayload struct {
	Token  string `json:"token"`
	UserID string `json:"userID"`
}

type Post struct {
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Subtitle  *string `json:"subtitle"`
	Content   string  `json:"content"`
	Slug      string  `json:"slug"`
	Author    *User   `json:"author"`
}

type User struct {
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	ID             string `json:"id"`
	HashedPassword string `json:"hashedPassword"`
	Email          string `json:"email"`
}
