package dto

import "mime/multipart"

type NewsDTO struct {
	Title  string               `json:"title" form:"title" binding:"required"`
	Author string               `json:"author" form:"author" binding:"required"`
	Status string               `json:"status" form:"status" binding:"required"`
	File   multipart.FileHeader `json:"file" form:"file" binding:"required"`
}
