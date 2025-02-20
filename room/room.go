package room

//import "lem-in/anthill"

type Room struct {
	Type string
	Name string
	Point_X int
	Point_y int
	Ants []uint
	Tunnel []map[string]int
}