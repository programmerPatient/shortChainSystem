package link

import (
	"github.com/laravelGo/core/database"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetByShortUrl(value string) (link Link) {
	database.DB.Where("short_url = ?", value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where(field+" = ?", value).Count(&count)
	return count > 0
}
