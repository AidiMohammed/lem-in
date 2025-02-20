package main

import (
	"lem-in/tools"
	"lem-in/anthill"
	"lem-in/room"
	"fmt"
)

func main(){

	myRoom := room.Room{Type: anthill.Commandes[anthill.INDEX_START]}
	myAnthill := anthill.Anthill{}

	err := myAnthill.SetRoom(myRoom)
	tools.HandelError(err,"")
	
	fmt.Println(myAnthill.Rooms)

	/*myAnthill,err := tools.LaodFileInput("./input01.txt")
	tools.HandelError(err,"")
	
	err = myAnthill.ValidateAnthill()

	if err != nil {
		fmt.Println(err)
	} else {
		myAnthill.ShowAnthill()
	}*/
} 