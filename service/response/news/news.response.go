package _news

import (
	"github.com/luvnyen/department-news-board/pkg/models"
)

type NewsResponse struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	File      string `json:"file"`
}

func NewNewsResponse(news models.News) NewsResponse {
	return NewsResponse{
		ID:        news.ID,
		Title:     news.Title,
		Author:    news.Author,
		Status:    news.Status,
		CreatedAt: news.CreatedAt.String(),
		File:      news.File,
	}
}

func NewNewsArrayResponse(news []models.News) []NewsResponse {
	var newsResponse []NewsResponse
	for _, news := range news {
		newsResponse = append(newsResponse, NewNewsResponse(news))
	}
	return newsResponse
}
