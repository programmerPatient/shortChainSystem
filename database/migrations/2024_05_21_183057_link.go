package migrations

import (
	"database/sql"

	"github.com/laravelGo/app/models"
	"github.com/laravelGo/core/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		ShortUrl    string `gorm:"type:varchar(255);not null;unique"`
		OriginalUrl string `gorm:"type:varchar(2083);unique;default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2024_05_21_183057_link", up, down)
}
