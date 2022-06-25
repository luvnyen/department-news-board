package repository

import (
	"time"

	"github.com/luvnyen/department-news-board/pkg/models"
	"gorm.io/gorm"
)

type NewsRepository interface {
	All() ([]models.News, error)
	FindByID(newsID string) (models.News, error)
	InsertNews(news models.News) (models.News, error)
	UpdateNews(news models.News) (models.News, error)
	ArchiveNews() ([]models.News, error)
	DeleteNews(newsID string) error
}

type newsConnection struct {
	connection *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsConnection{connection: db}
}

func (db *newsConnection) All() ([]models.News, error) {
	news := []models.News{}
	err := db.connection.Find(&news)
	if err.Error != nil {
		return news, err.Error
	}
	return news, nil
}

func (db *newsConnection) FindByID(newsID string) (models.News, error) {
	var news models.News
	err := db.connection.Where("id = ?", newsID).Take(&news)
	if err.Error != nil {
		return news, err.Error
	}
	return news, nil
}

func (db *newsConnection) InsertNews(news models.News) (models.News, error) {
	errSave := db.connection.Save(&news)
	if errSave.Error != nil {
		return news, errSave.Error
	}

	errFetch := db.connection.Find(&news)
	if errFetch.Error != nil {
		return news, errFetch.Error
	}

	return news, nil
}

func (db *newsConnection) UpdateNews(news models.News) (models.News, error) {
	err := db.connection.Save(&news)
	if err.Error != nil {
		return news, err.Error
	}

	errFetch := db.connection.Find(&news)
	if errFetch.Error != nil {
		return news, errFetch.Error
	}

	return news, nil
}

func (db *newsConnection) ArchiveNews() ([]models.News, error) {
	news := []models.News{}
	err := db.connection.Where("created_at < ?", time.Now().AddDate(0, -1, 0)).Find(&news)
	if err.Error != nil {
		return news, err.Error
	}

	for _, n := range news {
		n.Status = "Archived"
		err := db.connection.Save(&n)
		if err.Error != nil {
			return news, err.Error
		}
	}

	return news, nil
}

func (db *newsConnection) DeleteNews(newsID string) error {
	var news models.News
	err := db.connection.Where("id = ?", newsID).Take(&news).Delete(&news)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
