package anthill

import (
	"fmt"
	"errors"
	"lem-in/tools/common"
	"lem-in/room"
)

const INDEX_START = 0
const INDEX_END = 1
const INDEX_COMMENT = 2
const INDEX_MIDDLEROOM = 3

var Commandes = []string{"##start","##end","##comment","middleroom"}

type Anthill struct {
	Rooms []room.Room
}

func (this *Anthill) SetRoom(room room.Room) error {
	if !common.Contain(room.Type,Commandes) {
		return errors.New(fmt.Sprintf("Room (%v) invalid ",room.Type))
	}
	this.Rooms = append(this.Rooms,room)
	return nil
}

func (this *Anthill) AddRoomInAnthill(room room.Room)error {
	if room.Type != Commandes[0] || room.Type != Commandes[1] || room.Type != Commandes[3] {
		return errors.New(fmt.Sprintf("type room (%v) et invalide ",room.Type))
	}
	this.Rooms = append(this.Rooms,room)

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
	//
}