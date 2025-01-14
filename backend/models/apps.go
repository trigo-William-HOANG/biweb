package models

import "gorm.io/gorm"

type Apps struct {
	ID          uint   `json:"id" gorm:"primaryKey"` // Set as primary key
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Link        string `json:"link"`
}

func MigrateApps(db *gorm.DB) error {
	err := db.AutoMigrate(&Apps{})
	return err
}
