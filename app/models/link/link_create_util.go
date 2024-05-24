package link

import "github.com/laravelGo/core/database"

func Create(data *Link) (bool, error) {
	result := database.DB.Create(data)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
