// 第四单元实验代码
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	// 随机数种子初始化，可使每次运行结果发生改变，而非一直使用同一年份
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		year := rand.Intn(10000) + 1
		leap := year%4 == 0
		month := rand.Intn(12) + 1
		daysInMonth := 31
		switch month {
		case 2:
			daysInMonth = 28
			if leap {
				daysInMonth++
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		var day int
		day = rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}
