package main

import (
	"fmt"
	"math/rand"
	"time"
)

const earthMarsDistance = 62100000

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子，用于每次生成不同输出
	fmt.Println("太空航行公司\t\t\t飞行天数\t\t飞行类型\t\t价格（百万美元）")
	for i := 0; i < 10; i++ {
		companyId := rand.Intn(3)
		var company string
		switch companyId {
		case 0:
			company = "Virgin Galactic"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Space Adventures"
		}
		fmt.Print(company)
		switch companyId {
		case 0:
			fmt.Print("\t\t\t")
		case 1:
			fmt.Print("\t\t\t\t")
		case 2:
			fmt.Print("\t\t")
		}
		random := rand.Intn(15)
		spaceshipSpeed := random + 16
		daysToMars := earthMarsDistance / (spaceshipSpeed * 86400)
		fmt.Print(daysToMars, "\t\t\t")
		oneWayOrReturn := rand.Intn(2)
		switch oneWayOrReturn {
		case 0:
			fmt.Print("单程")
		case 1:
			fmt.Print("往返")
		}
		fmt.Print("\t\t\t")
		price := random + 36
		if oneWayOrReturn == 1 { // 如果是往返票，收取双倍费用
			price *= 2
		}
		fmt.Println(price)
	}
}
