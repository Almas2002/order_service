package model

type Point struct {
	Lat float32
	Lon float32
}

type Coordinate struct {
	From Point
	To   Point
}
