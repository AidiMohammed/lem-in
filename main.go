package main

import (
	"lem-in/tools"
	"lem-in/anthill"
	"lem-in/room"
)

func main(){

	ants := []uint{1,2,3,4}
	myAnthill := anthill.Anthill{}

	myRooms := []room.Room{}
	myRooms = append(myRooms,room.Room{Type: anthill.Commandes[anthill.INDEX_START],Name: "mohammed", Ants: ants})
	ants = []uint{}
	myRooms = append(myRooms,room.Room{Type: anthill.Commandes[anthill.INDEX_MIDDLEROOM],Name: "amine", Ants: ants})
	ants = []uint{}
	myRooms = append(myRooms,room.Room{Type: anthill.Commandes[anthill.INDEX_END],Name: "aidi", Ants: ants})

	err := myAnthill.InitAnthill(myRooms)
	if tools.HandelError(err,"") {
		return
	}
	myAnthill.ShowAnthill()

	/*myAnthill,err := tools.LaodFileInput("./input01.txt")
	tools.HandelError(err,"")
	
	err = myAnthill.ValidateAnthill()

	if err != nil {
		fmt.Println(err)
	} else {
		myAnthill.ShowAnthill()
	}*/
} 