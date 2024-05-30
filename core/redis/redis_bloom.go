package redis

import (
	"sync"

	redis_bloom_go "github.com/RedisBloom/redisbloom-go"
)

var once_bloom sync.Once

var bloom *redis_bloom_go.Client

// GetDefaultRedisBloomClient 初始化并返回一个默认的 RedisBloom 客户端实例。
// 该函数是线程安全的，使用了 sync.Once 确保客户端只被初始化一次。
// 返回值 *redis_bloom_go.Client: 表示初始化好的 RedisBloom 客户端实例。
func GetDefaultRedisBloomClient() *redis_bloom_go.Client {
	once_bloom.Do(func() {
		bloom = redis_bloom_go.NewClient("redis5.0:6379", "nohelp", nil)
		// 通过配置信息初始化 RedisBloom 客户端。
		// bloom = GetRedisBloomClient(config.GetString("bloom_redis.name"), config.GetString("bloom_redis.host"), config.GetString("bloom_redis.port"), nil)
	})
	return bloom
}

// GetRedisBloomClient 创建并返回一个 Redis Bloom 客户端实例。
// 该函数接收以下参数：
// - name: 客户端的名称，用于标识不同的客户端实例。
// - host: Redis 服务器的主机名或IP地址。
// - port: Redis 服务器的端口号。
// - authPass: 连接Redis服务器时使用的认证密码，如果不需要认证，则可以传入nil。
// 返回值为一个已初始化的 Redis Bloom 客户端实例指针。
func GetRedisBloomClient(name, host, port string, authPass *string) *redis_bloom_go.Client {
	// 从配置文件中读取 Redis 服务器的主机名和端口号，并格式化为地址字符串。
	// // address := fmt.Sprintf("%v:%v", host, port)
	// // 使用配置的地址、客户端名称和认证密码创建 Redis Bloom 客户端实例。
	// // 使用布隆过滤器操作 Redis
	// bf := redis.NewBloomFilter("myfilter", 1000000, 0.01)
	// bf.Reserve("foo", "bar")
	// exists, err := bf.Exists("foo")
	// if err != nil {
	// 	panic(err)
	// }
	return nil
}
