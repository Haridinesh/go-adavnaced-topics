package lookup

import (
	"blogpost/models"

	"gorm.io/gorm"
)

func (*Formethod) Lookup_001(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Logincredentials{}, &models.Posts{}, &models.Categories{}, &models.Comments{}, &Lookups{}); err != nil {
		return
	}
}
