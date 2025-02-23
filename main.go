package main

import (
	//"fmt"
	"lem-in/tools"
	"lem-in/tools/common"
	"lem-in/room"
	"lem-in/anthill"
)

func main(){

	myAnthill := anthill.Anthill{}

	typeRoom := common.Commandes[common.INDEX_START]
	nameRoom := "Mohammed"
	point_x := 3
	point_y := 6
	ants := []uint{1}

	myRooms := []room.Room{}
	NewRoom,err := room.MakeRoom(typeRoom,nameRoom,point_x,point_y,ants)
	if tools.HandelError(err,"") {
		//fmt.Println(err)
		return
	}
	myRooms = append(myRooms,NewRoom)

	typeRoom = common.Commandes[common.INDEX_MIDDLEROOM]
	nameRoom = "Amine"
	point_x = 3
	point_y = 6
	ants = []uint{}

	NewRoom = room.Room{}
	NewRoom,err = room.MakeRoom(typeRoom,nameRoom,point_x,point_y,ants)
	if tools.HandelError(err,"") {
		return
	}
	myRooms = append(myRooms,NewRoom)

	typeRoom = common.Commandes[common.INDEX_END]
	nameRoom = "Fidal"
	point_x = 8
	point_y = 5
	ants = []uint{}

	NewRoom = room.Room{}
	NewRoom,err = room.MakeRoom(typeRoom,nameRoom,point_x,point_y,ants)
	if tools.HandelError(err,"") {
		return
	}
	myRooms = append(myRooms,NewRoom)

	err = myAnthill.InitAnthill(myRooms)
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