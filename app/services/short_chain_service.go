package services

import (
	"sync"

	"github.com/laravelGo/app/models/link"
)

type ShortChainService struct{}

var sc_once sync.Once

var sc_s *ShortChainService

func GetShortChainService() *ShortChainService {
	sc_once.Do(func() {
		sc_s = &ShortChainService{}
	})
	return sc_s
}

// 根据短链接获取原始链接
func (s *ShortChainService) GetOriginalUrlByShortUrl(shortUrl string) string {
	link := link.GetBy("short_url", shortUrl)
	return link.OriginalUrl
}
