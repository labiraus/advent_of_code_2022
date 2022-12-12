package main

type world interface {
	tile(x, y int) *tile
}
type layout struct {
	w [][]*tile
}

func (w *layout) tile(x, y int) *tile {
	if x >= 0 && x < len(w.w) && y >= 0 && y < len(w.w[x]) {
		return w.w[x][y]
	}
	return nil
}
