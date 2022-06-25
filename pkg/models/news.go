package models

import "time"

type News struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Author    string    `gorm:"type:varchar(255)" json:"author"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	File      string    `gorm:"type:varchar(255)" json:"file"`
}
