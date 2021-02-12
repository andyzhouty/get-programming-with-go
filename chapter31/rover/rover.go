package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

// RoverDriver用于驱动一台在火星表面行进的探测器
type RoverDriver struct {
	commandc chan command
}

// driver 负责驱动探测器。这个方法放在goroutine中运行
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 0
	for {
		select {
		case c := <- r.commandc:
			switch c {
			case right: // 右转
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case stop:
				fmt.Printf("Stopping\n")
				speed = 0
			case start:
				fmt.Printf("Starting\n")
				speed = 1
			}
			log.Printf("new direction %v; current speed %v", direction, speed)
		case <- nextMove:
			pos = pos.Add(direction.Mul(speed))
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// Left会讲探测器转向左方（逆时针90°）
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right会将探测器转向右方
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	right = command(0)
	left = command(1)
	start = command(2)
	stop = command(3)
)

func main() {
	r := NewRoverDriver()
	r.Start()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(time.Second)
}
