package anthill

import (
	"fmt"
	//"errors"
	"lem-in/room"
)

var Flags = []string{"##start","##end","##comment"}

type Anthill struct {
	Rooms []room.Room
}

func (this *Anthill) AddRoomInAnthill(room room.Room) error{
	return nil
}

func (this Anthill) ValidateAnthill() error {
	return nil
}

func (this Anthill) ShowAnthill() {

	fmt.Println("====================================")
	fmt.Println("  📜 Information  de la Foumilière  ")
	fmt.Println("====================================")

	
	fmt.Println("====================================")
}

func (this Anthill) GetRoomType(typeRoom string) room.Room {
	myRoom := room.Room{}
	return myRoom
}

func (this *Anthill) AddAntsInRoom(countAnt int,typeRoom string, nameRoom string) error {
	return nil
}

func (this *Anthill) initAnthill(){
	this.
}