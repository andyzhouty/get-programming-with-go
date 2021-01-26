package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	for i, c := range cipherText {
		var c2 int
		c2 = int(c) - int(keyword[i % 6]) // 取得两个字符相减的整数值
		c2 += 26 // c2自增26，此时c2可能大于26
		c2 = c2 % 26 + 65 // 使用c2除以26的值来获取c2在字母顺序中的排位在加上65来表示ASCII码
		fmt.Printf("%c", rune(c2)) // 输出
	}
}
