package main

import (
	"lem-in/tools"
	"lem-in/anthill"
	"lem-in/room"
	"fmt"
)

func main(){

	ants := []uint{1,2,3,4}
	myAnthill := anthill.Anthill{}
	var err error

	myRooms := []room.Room{}
	myRooms = append(myRooms,room.Room{Type: anthill.Commandes[anthill.INDEX_START],Name: "mohammed", Ants: ants})
	myRooms = append(myRooms,room.Room{Type: anthill.Commandes[anthill.INDEX_END],Name: "amine", Ants: ants})

	for _,room := range myRooms {
		err = myAnthill.AddRoomInAnthill(room)
		if err != nil {
			break
		}
	}
	if tools.HandelError(err,"") {
		return
	}
	
	fmt.Println(myAnthill.Rooms)
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