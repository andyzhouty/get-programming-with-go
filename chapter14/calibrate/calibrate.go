package main

import "fmt"

type kelvin float64

// sensor函数类型
type sensor func() kelvin

func realSensor() kelvin {
	return 0 // 待办事项：实现真正传感器
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}
func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
