package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type item int

const (
	unscanned item = iota
	potential
	scanned
	sensorItem
	beaconItem
)

type Dataset struct {
	MinX     int
	MinY     int
	MaxX     int
	MaxY     int
	Coverage map[int]map[int]item
	Sensors  []*Sensor
}

type Sensor struct {
	Position  Coordinate
	Beacon    Coordinate
	Radius    int
	MaxRadius int
}

type Coordinate struct {
	X int
	Y int
}

func main() {
	f, err := os.Open("day15/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make(chan string)
	go func() {
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	d := buildDataset()
	fmt.Println("building")
	d.build(lines)
	d.MinX = 0
	d.MinY = 0
	d.MaxX = 4000000
	d.MaxY = 4000000
	fmt.Println("plotting")
	x, y := d.plot(4000000, 4000000)

	// fmt.Println(d.eval(2000000))
	fmt.Println(x*4000000 + y)
}

func buildDataset() Dataset {
	d := Dataset{Coverage: make(map[int]map[int]item), Sensors: make([]*Sensor, 0)}
	return d
}

func (d *Dataset) build(lines <-chan string) {
	r, _ := regexp.Compile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for line := range lines {
		line = strings.TrimSpace(line)
		sensor := buildSensor(r.FindStringSubmatch(line))

		if d.MinX > sensor.Position.X-sensor.Radius {
			d.MinX = sensor.Position.X - sensor.Radius
		}
		if d.MaxX < sensor.Position.X+sensor.Radius {
			d.MaxX = sensor.Position.X + sensor.Radius
		}
		if d.MinY > sensor.Position.Y-sensor.Radius {
			d.MinY = sensor.Position.Y - sensor.Radius
		}
		if d.MaxY < sensor.Position.Y+sensor.Radius {
			d.MaxY = sensor.Position.Y + sensor.Radius
		}
		d.setVal(sensor.Position.X, sensor.Position.Y, sensorItem)
		d.setVal(sensor.Beacon.X, sensor.Beacon.Y, beaconItem)
		d.Sensors = append(d.Sensors, &sensor)
	}
}

func (d *Dataset) plot(maxX, maxY int) (int, int) {
	for x := 0; x <= maxX; x++ {
		// if x%10000 == 0 {
		// 	fmt.Println("scanning", x)
		// }
		for y := 0; y <= maxY; y++ {
			if _, ok := d.Coverage[x]; ok && d.Coverage[x][y] == beaconItem {
				continue
			}
			found := true
			for _, sensor := range d.Sensors {
				radius := sensor.radius(Coordinate{X: x, Y: y})
				if radius <= sensor.Radius {
					found = false
					y += sensor.Radius - radius
					break
				}
			}
			if found {
				return x, y
			}
		}
	}
	return 0, 0
}

func (d *Dataset) compareSensors() {
	for i, sensor := range d.Sensors {
		fmt.Println("comparing", i)
		for j := i + 1; j < len(d.Sensors); j++ {
			radius := sensor.radius(d.Sensors[j].Beacon)
			if radius < sensor.MaxRadius {
				sensor.MaxRadius = radius
			}
			if radius < d.Sensors[j].MaxRadius {
				d.Sensors[j].MaxRadius = radius
			}
		}
		for x := sensor.Position.X - sensor.MaxRadius; x <= sensor.Position.X+sensor.MaxRadius; x++ {
			for y := sensor.Position.Y - sensor.MaxRadius; y <= sensor.Position.Y+sensor.MaxRadius; y++ {
				radius := sensor.radius(Coordinate{X: x, Y: y})
				if radius <= sensor.MaxRadius && radius > sensor.Radius {
					d.setVal(x, y, potential)
				} else if radius <= sensor.Radius {
					d.setVal(x, y, scanned)
				}
			}
		}
	}
}

func buildSensor(matches []string) Sensor {
	var err error
	sensor := Sensor{MaxRadius: 1000000}
	sensor.Position.X, err = strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	sensor.Position.Y, err = strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	sensor.Beacon.X, err = strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}
	sensor.Beacon.Y, err = strconv.Atoi(matches[4])
	if err != nil {
		panic(err)
	}

	sensor.Radius = sensor.radius(sensor.Beacon)
	return sensor
}

func (s *Sensor) radius(coord Coordinate) int {
	return int(math.Abs(float64(s.Position.X)-float64(coord.X))) + int(math.Abs(float64(s.Position.Y)-float64(coord.Y)))
}

func (d *Dataset) eval(row int) int {
	fmt.Println("evaulating")
	sum := 0
	for x := d.MinX; x <= d.MaxX; x++ {
		if _, ok := d.Coverage[x]; ok && d.Coverage[x][row] == beaconItem {
			continue
		}

		for _, sensor := range d.Sensors {
			radius := sensor.radius(Coordinate{X: x, Y: row})
			if radius <= sensor.Radius {
				sum++
				break
			}
		}

	}

	// for _, col := range d.Coverage {
	// 	switch col[row] {
	// 	case scanned:
	// 		sum++
	// 	case sensorItem:
	// 		sum++
	// 	}
	// }

	return sum
}

func (d Dataset) String() string {
	out := "  "
	for x := d.MinX; x <= d.MaxX; x++ {
		if x%5 == 0 && x > 0 {
			out += strconv.Itoa(x % 100 / 10)
		} else {
			out += " "
		}
	}
	out += "\n  "
	for x := d.MinX; x <= d.MaxX; x++ {
		if x%5 == 0 && x > 0 {
			out += strconv.Itoa(x % 10 / 10)
		} else {
			out += " "
		}
	}
	out += "\n"
	for y := d.MinY; y <= d.MaxY; y++ {
		out += fmt.Sprintf("%2d", y)
		for x := d.MinX; x <= d.MaxX; x++ {
			if _, ok := d.Coverage[x]; !ok {
				d.Coverage[x] = make(map[int]item)
			}
			val, ok := d.Coverage[x][y]
			if !ok {
				out += "."
				continue
			}
			switch val {
			case beaconItem:
				out += "B"
			case sensorItem:
				out += "S"
			case scanned:
				out += "#"
			case unscanned:
				out += "."
			case potential:
				out += "?"
			}

		}
		out += "\n"
	}
	return out
}

func (d *Dataset) setVal(x, y int, val item) {
	if _, ok := d.Coverage[x]; !ok {
		d.Coverage[x] = make(map[int]item)
	}
	currentVal := d.Coverage[x][y]
	if val > currentVal {
		d.Coverage[x][y] = val
	}
}
