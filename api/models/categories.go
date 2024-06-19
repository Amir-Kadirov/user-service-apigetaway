package models

import (
	"time"
)

type GetAllCategoryRequest struct{
	Offset int64 
	Limit int64
	Search string 
} 

type GetAllCategoryResponse struct {
	Count int64
	Categories []Categories
}

type Categories struct{
	Id string
	Slug string 
	NameUz string
	NameRu string
	NameEn string
	Active bool
	OrderNo int64
	ParentId string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}