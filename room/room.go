package room

type Room struct {
	Type string
	Name string
	Point_X int
	Point_y int
	Ants []int
	Tunnel []map[string]int
}