package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Status      string    `json:"status"`
}

type CreatePostRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

type UpdatePostRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}
