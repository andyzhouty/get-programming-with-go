package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

// MarsGrid 网格用于表示火星的某些表面
// 他可能会被多个不同的goroutine并发使用
type MarsGrid struct {
	mu     sync.Mutex
	bounds image.Rectangle
	cells  [][]cell
}

// Occupier 用于表示网格中一个已被占据的单元格
// 它可能会被多个不同的goroutine并发使用
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

type cell struct {
	groundData SensorData
	occupier   *Occupier
}

// RoverDriver用于驱动一台在火星表面行进的探测器
type RoverDriver struct {
	commandc chan command
	occupier *Occupier
	name     string
	radio    *Radio
}

type command int

type Message struct {
	Pos       image.Point
	LifeSigns int
	Rover     string
}

type Radio struct {
	fromRover chan Message
}

type SensorData struct {
	LifeSigns int
}

func earthReceiver(messages chan []Message) {
	for {
		time.Sleep(dayLength - receiveTimePerDay)
		receiveMarsMessages(messages)
	}
}

func receiveMarsMessages(messages chan []Message) {
	finished := time.After(receiveTimePerDay)
	for {
		select {
		case <-finished:
			return
		case ms := <-messages:
			for _, m := range ms {
				log.Printf("earth received report of life sign level %d from %s at %v", m.LifeSigns, m.Rover, m.Pos)
			}
		}
	}
}

func (r *Radio) SendToEarth(m Message) {
	r.fromRover <- m
}

func NewRadio(toEarth chan []Message) *Radio {
	r := &Radio{
		fromRover: make(chan Message),
	}
	go r.run(toEarth)
	return r
}

func (r *Radio) run(toEarth chan []Message) {
	var buffered []Message
	for {
		toEarth1 := toEarth
		if len(buffered) == 0 {
			toEarth1 = nil
		}
		select {
		case m := <-r.fromRover:
			buffered = append(buffered, m)
		case toEarth1 <- buffered:
			buffered = nil
		}
	}
}

// Occupy 占据网格中给定坐标点上的单元格
// 它在单元格已经被占据或者坐标点不在网格范围内时返回nil
// 否则它将返回一个值，该值可以用于将单元格一直网格的其他位置
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.getCell(p)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

// 获取一个单元格
func (g *MarsGrid) getCell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}

// MoveTo 会尝试将给定的Occupier 移至网格中的其他单元格。然后报告移动是否成功
// 如果移动超出网格范围或者移动的目标单元格已被占据，那么移动将会失败
// 在移动失败的情况下，Occupier将会继续留在原来的单元格
func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	newCell := o.grid.getCell(p)
	if newCell == nil || newCell.occupier != nil {
		return false
	}
	o.grid.getCell(o.pos).occupier = nil
	o.pos = p
	return true
}

func (o *Occupier) Sense() SensorData {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	return o.grid.getCell(o.pos).groundData
}

func (r *RoverDriver) checkForLife() {
	sensorData := r.occupier.Sense()
	if sensorData.LifeSigns < 900 {
		return
	}
	r.radio.SendToEarth(Message{
		Pos:       r.occupier.pos,
		LifeSigns: sensorData.LifeSigns,
		Rover:     r.name,
	})
}

// driver 负责驱动探测器。这个方法放在goroutine中运行
func (r *RoverDriver) drive() {
	log.Printf("%s initial position %v", r.name, r.occupier.pos)
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
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
			}
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.pos.Add(direction)
			if r.occupier.MoveTo(newPos) {
				log.Printf("%s moved to %v", r.name, newPos)
				r.checkForLife()
				break
			}
			log.Printf("%s blocked trying to move from %v to %v", r.name, r.occupier.pos, newPos)
			dir := rand.Intn(3) + 1
			for i := 0; i < dir; i++ {
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s new random direction %v", r.name, direction)
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

func NewRoverDriver(name string, grid *MarsGrid, marsToEarth chan []Message) *RoverDriver {
	var o *Occupier
	// 尝试一个随机点直到我们找到一个没有被占据的点
	var startPoint image.Point
	for o == nil {
		startPoint = image.Point{X: rand.Intn(x), Y: rand.Intn(y)}
		o = grid.Occupy(startPoint)
	}
	r := &RoverDriver{
		commandc: make(chan command),
		occupier: o,
		name:     name,
		radio:    NewRadio(marsToEarth),
	}
	go r.drive()
	return r
}

func NewMarsGrid(point image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{Max: point},
		cells:  make([][]cell, y),
	}
	for i := range grid.cells {
		grid.cells[i] = make([]cell, x)
		for j := range grid.cells[i] {
			cell := &grid.cells[i][j]
			cell.groundData.LifeSigns = rand.Intn(1000)
		}
	}
	return grid
}

const (
	right             = command(0)
	left              = command(1)
	x                 = 50
	y                 = 60
	dayLength         = 24 * time.Second
	receiveTimePerDay = 2 * time.Second
)

func main() {
	marsToEarth := make(chan []Message)
	go earthReceiver(marsToEarth)

	gridSize := image.Point{X: x, Y: y}
	grid := NewMarsGrid(gridSize)
	rover := make([]*RoverDriver, 5)
	for i := range rover {
		rover[i] = NewRoverDriver(fmt.Sprint("rover#", i), grid, marsToEarth)
	}
	time.Sleep(60 * time.Second)
}
