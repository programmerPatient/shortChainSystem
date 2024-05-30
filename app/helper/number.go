package helper

import (
	"math/rand"
)

// 数字转化为
func ToBase64(num uint64, length int) string {
	const base = 64
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz+/"

	//添加随机性，防止生成的字符串有序可查
	charsArray := []rune(chars)
	rand.Shuffle(len(charsArray), func(i, j int) {
		charsArray[i], charsArray[j] = charsArray[j], charsArray[i]
	})
	shuffledChars := string(charsArray)
	if num == 0 {
		return "0"
	}
	return_str := ""
	for num > 0 {
		return_str = string(shuffledChars[num%base]) + return_str
		num /= base
		length--
	}
	for i := 0; i < length; i++ {
		//添加随机性，防止生成的字符串有序可查
		charsArray := []rune(chars)
		rand.Shuffle(len(charsArray), func(i, j int) {
			charsArray[i], charsArray[j] = charsArray[j], charsArray[i]
		})
		shuffledChars := string(charsArray)
		return_str = string(shuffledChars[0]) + return_str
	}
	return return_str
}
