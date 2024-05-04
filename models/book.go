package models

import "time"

type Book struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `gorm:"type:varchar(100)" json:"title"`
	AuthorID  int    `gorm:"type:int" json:"author_id"`
	Author    Author
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookResource struct {
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	Author    AuthorResource `json:"author"`
	CreatedAt time.Time      `json:"created_at"`
}
