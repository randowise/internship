package migrations

import (
	"example/APIIIDO0/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Pro{})
}
