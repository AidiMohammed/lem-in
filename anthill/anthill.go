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

func (this *Anthill) AddRoomInAnthill(room room.Room) error {

	fmt.Println(fmt.Sprintf("🏗️ Room ( %v ) creation is in progress, please wait...",room.Name))

	err := false
	errMessage := ""

	for _,thisRoom := range this.Rooms {
		if room.Type == Commandes[INDEX_START] && thisRoom.Type == Commandes[INDEX_START] {
			err = true
			errMessage = fmt.Sprintf("Error ❌ : (%v)\nAtnhill can't have two rooms Start !",room.Name)
		}
	}

	if !common.Contain(room.Type,Commandes) {
		err = true
		errMessage = fmt.Sprintf("Error ❌ : \nRoom (%v) invalid ",room.Type)
	} else if room.Name == "" {
		err = true
		errMessage = ("Error ❌ : \nThe room must have a name")
	} else if room.Type == Commandes[INDEX_MIDDLEROOM] && len(room.Ants) > 0 {
		err = true
		errMessage = ("Error ❌ : \nIn middleroom should only contain one ant")
	}

	if err {
		return errors.New(common.ColorString(common.INDEX_C_RED,errMessage))
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

	for _,room := range this.Rooms {
		fmt.Println(fmt.Sprintf("Type Room : %v",room.Type))
		fmt.Println(fmt.Sprintf("Name room : %v",room.Name))
		fmt.Println(fmt.Sprintf("Number ants in room : %v",len(room.Ants)))
		fmt.Println("====================================")

	}
	
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