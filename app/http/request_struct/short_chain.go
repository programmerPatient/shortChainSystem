package requeststruct

// 创建短链接接口的参数处理结构体
type ShortChainCreate struct {
	OriginalUrl string `json:"original_url" binding:"required"` // 原始链接
}
