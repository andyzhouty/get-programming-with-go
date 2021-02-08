package main

import "fmt"

type report struct {
	sol
	location
	temperature
}

type sol int
type location struct {
	lat, long float64
}

type celsius float64

type temperature struct {
	high, low celsius
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	// 待办事项：复杂的距离计算
	return 5
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	report := report{sol: 15}
	d := report.days(1446)
	fmt.Println(d)
}
