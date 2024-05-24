package helper

// 数字转化为
func ToBase64(num uint64, length int) string {
	const base = 64
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz+/"

	if num == 0 {
		return "0"
	}
	return_str := ""
	for num > 0 {
		return_str = string(chars[num%base]) + return_str
		num /= base
		length--
	}
	for i := 0; i < length; i++ {
		return_str = string(chars[0]) + return_str
	}
	return return_str
}
