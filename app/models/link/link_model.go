package link

import "github.com/laravelGo/app/models"

type Link struct {
	models.BaseModel
	// 替换成具体的结构
	ShortUrl    string `gorm:"column:short_url;" json:"short_url,omitempty"`       //短链接
	OriginalUrl string `gorm:"column:original_url;" json:"original_url,omitempty"` //真实链接
	models.CommonTimestampsField
}

// 自定义表名
func (Link) TableName() string {
	return "links"
}
