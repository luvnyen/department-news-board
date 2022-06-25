package service

import (
	"time"

	"github.com/luvnyen/department-news-board/pkg/dto"
	"github.com/luvnyen/department-news-board/pkg/models"
	"github.com/luvnyen/department-news-board/repository"
	_news "github.com/luvnyen/department-news-board/service/response/news"
)

type NewsService interface {
	All() (*[]_news.NewsResponse, error)
	FindByID(newsID string) (*_news.NewsResponse, error)
	InsertNews(newsRequest dto.NewsDTO, newFileName string) (*_news.NewsResponse, error)
	UpdateNews(newsID string, newsRequest dto.NewsDTO, newFileName string) (*_news.NewsResponse, error)
	ArchiveNews() (*[]_news.NewsResponse, error)
	DeleteNews(newsID string) error
}

type newsService struct {
	newsRepository repository.NewsRepository
}

func NewNewsService(newsRepository repository.NewsRepository) NewsService {
	return &newsService{newsRepository: newsRepository}
}

func (service *newsService) All() (*[]_news.NewsResponse, error) {
	news, err := service.newsRepository.All()
	if err != nil {
		return nil, err
	}
	news_res := _news.NewNewsArrayResponse(news)
	return &news_res, nil
}

func (service *newsService) FindByID(newsID string) (*_news.NewsResponse, error) {
	news, err := service.newsRepository.FindByID(newsID)
	if err != nil {
		return nil, err
	}
	news_res := _news.NewNewsResponse(news)
	return &news_res, nil
}

func (service *newsService) InsertNews(newsRequest dto.NewsDTO, newFileName string) (*_news.NewsResponse, error) {
	news := models.News{}
	news.Title = newsRequest.Title
	news.Author = newsRequest.Author
	news.Status = newsRequest.Status
	news.CreatedAt = time.Now()
	news.File = newFileName

	news, err := service.newsRepository.InsertNews(news)
	if err != nil {
		return nil, err
	}

	news_res := _news.NewNewsResponse(news)
	return &news_res, nil
}

func (service *newsService) UpdateNews(newsID string, newsRequest dto.NewsDTO, newFileName string) (*_news.NewsResponse, error) {
	news, err := service.newsRepository.FindByID(newsID)
	if err != nil {
		return nil, err
	}

	news.Title = newsRequest.Title
	news.Author = newsRequest.Author
	news.Status = newsRequest.Status
	news.File = newFileName

	updatedNews, err := service.newsRepository.UpdateNews(news)
	if err != nil {
		return nil, err
	}

	res := _news.NewNewsResponse(updatedNews)
	return &res, nil
}

func (service *newsService) ArchiveNews() (*[]_news.NewsResponse, error) {
	news, err := service.newsRepository.ArchiveNews()
	if err != nil {
		return nil, err
	}

	news_res := _news.NewNewsArrayResponse(news)
	return &news_res, nil
}

func (service *newsService) DeleteNews(newsID string) error {
	err := service.newsRepository.DeleteNews(newsID)
	if err != nil {
		return err
	}

	return nil
}
