package anthill

import (
	"fmt"
	"errors"
	"lem-in/tools/common"
	"lem-in/room"
)

type Anthill struct {
	Rooms []room.Room
	Tunnels []map[*room.Room][]*room.Room

}

func (this *Anthill) InitAnthill(rooms []room.Room) error {
	
	common.BoxString("⏳ Init the Anthill in progress, please wait...")

	if len(this.Rooms) > 0 {
		this.Rooms = []room.Room{}
	}
	for _,room := range rooms {
		this.Rooms = append(this.Rooms,room)
	}
	if !this.isValidAnthill() {
		return errors.New(common.ColorString(common.INDEX_C_RED,"Error ❌ : \nThe anthill is not valid"))
	}	
	fmt.Println(common.ColorString(common.INDEX_C_GREEN,"✅ Your Anthill had been initialized succesfully !"))
	return nil 
}

func (this *Anthill) addNewRoomInAnthill(room room.Room) {
	this.Rooms = append(this.Rooms,room)
}

func (this Anthill) isValidAnthill() bool {
	startRoom := false
	endRoom := false
	middleRoom := false

	for _,room := range this.Rooms {
		if room.Type == common.Commandes[common.INDEX_START] {
			startRoom = true
		} else if room.Type == common.Commandes[common.INDEX_MIDDLEROOM] {
			middleRoom = true
		} else if room.Type == common.Commandes[common.INDEX_END] {
			endRoom = true
		}
	}
	if startRoom && middleRoom && endRoom {
		return true
	}

	return false
}

func (this Anthill) ShowAnthill() {

	common.BoxString("Information  de la Foumilière")

	for _,room := range this.Rooms {
		fmt.Println(fmt.Sprintf("Type Room : %v",room.Type))
		fmt.Println(fmt.Sprintf("Name room : %v",room.Name))
		fmt.Println(fmt.Sprintf("Number ants in room : %v",len(room.Ants)))
		fmt.Println("====================================")

	}
	
}

func (this *Anthill) MakeNodes(kayRom *room.Room, valueRoom *room.Room) {
	fmt.Println("Make the Nodes : ")
	fmt.Println(kayRom)
	fmt.Println(valueRoom)
}

func (this Anthill) GetRoomByName(nameRoom string) (*room.Room,error) {
	for _,room := range this.Rooms {
		if room.Name == nameRoom {
			return &room,nil
		}
	}
	emptyRoom := room.Room{}
	errMessage := fmt.Sprintf("Error ❌ : \nThe room %v is not found in the anthill ",nameRoom)
	return &emptyRoom,errors.New(common.ColorString(common.INDEX_C_RED,errMessage)) 
}