package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal方法会将DMS格式的坐标转换为十进制格式
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	// 布莱德伯利着陆点：南纬4°35'22.2", 东经137°26'30.1"
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
