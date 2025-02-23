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

func (this *Anthill) InitAnthill(rooms []room.Room) error {
	
	common.BoxString("===== ⏳ Init the Anthill in progress, please wait... =====")

	if len(this.Rooms) > 0 {
		this.Rooms = []room.Room{}
	}
	for _,room := range rooms {
		err := this.addNewRoomInAnthill(room)
		if err != nil {
			return err
		}
	}
	if !this.isValidAnthill() {
		return errors.New(common.ColorString(common.INDEX_C_RED,"Error ❌ : \nThe anthill is not valid"))
	}	
	fmt.Println(common.ColorString(common.INDEX_C_GREEN,"✅ Your Anthill had been initialized succesfully !"))
	return nil 
}

func (this *Anthill) addNewRoomInAnthill(room room.Room) error {

	fmt.Printf("⏳ Room ( %v ) creation is in progress, please wait...\n",room.Name)

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
		errMessage = fmt.Sprintf("Error ❌ : \nThe room (%v) invalid, unknown command !",room.Type)
	} else if room.Name == "" {
		err = true
		errMessage = ("Error ❌ : \nThe room must have a name !")
	} else if room.Type == Commandes[INDEX_MIDDLEROOM] && len(room.Ants) > 0 {
		err = true
		errMessage = ("Error ❌ : \nThe middle room should not contain ants !")
	} else if room.Type == Commandes[INDEX_END] && len(room.Ants) > 0 {
		err = true
		errMessage = ("Error ❌ : \nThe end room shpuld not contain ants !")
	}

	if err {
		return errors.New(common.ColorString(common.INDEX_C_RED,errMessage))
	} 
	this.Rooms = append(this.Rooms,room)
	return nil
}

func (this Anthill) isValidAnthill() bool {
	startRoom := false
	endRoom := false
	middleRoom := false

	for _,room := range this.Rooms {
		if room.Type == Commandes[INDEX_START] {
			startRoom = true
		} else if room.Type == Commandes[INDEX_MIDDLEROOM] {
			middleRoom = true
		} else if room.Type == Commandes[INDEX_END] {
			endRoom = true
		}
	}
	if startRoom && middleRoom && endRoom {
		return true
	}

	return false
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