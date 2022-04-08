package models

import (
	"time"

	"gorm.io/gorm"
)

//earnings fazer com o float64
type Pro struct {
	ID             string          `json:"id" `
	Name           string          `json:"name"`
	UserName       string          `json:"user_name"`
	Description    string          `json:"description"`
	Mouse          string          `json:"mouse"`
	SettingsInGame string          `json:"settings_in_game"`
	Keyboard       string          `json:"keyboard"`
	SocialMedia    string          `json:"social_media"`
	Earnings       string          `json:"earnings"`
	CreatedAt      *time.Time      `json:"created_at"`
	UpdatedAt      *time.Time      `json:"updated_at"`
	DeletedAt      *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
