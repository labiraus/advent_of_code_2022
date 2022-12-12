package main

import (
	"github.com/beefsack/go-astar"
)

type tile struct {
	X       int
	Y       int
	height  int
	path    bool
	touched bool
	layout  world
}

// PathEstimatedCost implements astar.Pather
func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

// PathNeighborCost implements astar.Pather
func (t *tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t *tile) PathNeighbors() []astar.Pather {
	tiles := make([]astar.Pather, 0)
	if neighbour := t.layout.tile(t.X-1, t.Y); neighbour != nil && t.compareHeights(neighbour) {
		tiles = append(tiles, neighbour)
		neighbour.touched = true
	}
	if neighbour := t.layout.tile(t.X+1, t.Y); neighbour != nil && t.compareHeights(neighbour) {
		tiles = append(tiles, neighbour)
		neighbour.touched = true
	}
	if neighbour := t.layout.tile(t.X, t.Y-1); neighbour != nil && t.compareHeights(neighbour) {
		tiles = append(tiles, neighbour)
		neighbour.touched = true
	}
	if neighbour := t.layout.tile(t.X, t.Y+1); neighbour != nil && t.compareHeights(neighbour) {
		tiles = append(tiles, neighbour)
		neighbour.touched = true
	}
	return tiles
}

func (t *tile) compareHeights(to *tile) bool {
	return to.height-t.height <= 1
}
