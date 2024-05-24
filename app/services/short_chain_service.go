package services

import (
	"sync"

	"github.com/laravelGo/app/helper"
	"github.com/laravelGo/app/models/link"
	"github.com/laravelGo/core/redis"
)

type ShortChainService struct{}

var sc_once sync.Once

var sc_s *ShortChainService

// GetShortChainService 是一个用于获取ShortChainService实例的函数。
// 该函数使用了单例模式来确保整个应用中只存在一个ShortChainService实例。
// 返回值 *ShortChainService: 返回ShortChainService的实例指针。
func GetShortChainService() *ShortChainService {
	// 使用sync.Once确保只初始化一次ShortChainService实例
	sc_once.Do(func() {
		sc_s = &ShortChainService{}
	})
	return sc_s
}

// 根据短链接获取原始链接
// GetOriginalUrlByShortUrl 通过短链接获取原始长链接
// 参数:
//
//	shortUrl string - 短链接的字符串表示
//
// 返回值:
//
//	string - 原始长链接的字符串表示
func (s *ShortChainService) GetOriginalUrlByShortUrl(shortUrl string) string {
	// 根据短链接查询数据库，获取对应记录
	link := link.GetByShortUrl(shortUrl)
	// 返回查询到的记录中的原始长链接
	return link.OriginalUrl
}

// CreateShortUrl 用于创建一个短链接。
// 参数 OriginalUrl 是需要被缩短的原始URL字符串。
// 返回值 bool 表示短链接是否创建成功。
// 返回值 error 表示在创建过程中遇到的任何错误。
func (s *ShortChainService) CreateShortUrl(OriginalUrl string) (*link.Link, error) {
	// 通过Redis的Incr方法自增索引，为每个原始URL生成一个唯一的ID
	index := redis.Redis.Incr("short_chain")
	// 将索引值转换为6位的Base64字符串作为短链接的地址
	str := helper.ToBase64(uint64(index), 6)
	data := &link.Link{
		ShortUrl:    str,
		OriginalUrl: OriginalUrl,
	}
	// 调用link.Create方法创建短链接，并返回结果
	_, err := link.Create(data)
	return data, err
}
