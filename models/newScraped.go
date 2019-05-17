package models

import (
	"time"
)

type NewScraped struct {
	Page      int       `json:"page"`
	FullPage  bool      `json:"full_page"`
	Headline  string    `json:"headline"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	NewsPaper string    `json:"newspaper"`
	ScraperID string    `json:"scraper_id" bson:"scraper_id"`
	ID        string    `json:"id"`
}
