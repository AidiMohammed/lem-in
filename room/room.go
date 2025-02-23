package room

import (
	"lem-in/tools/common"
	"fmt"
	"errors"
)

type Room struct {
	Type string
	Name string
	Point_X int
	Point_y int
	Ants []uint
}

func MakeRoom(typeRoom string, nameRoom string, point_x int, point_y int, ants []uint) (Room,error) {
	var myRoom Room

	fmt.Printf("⏳ Room ( %v ) creation is in progress, please wait...",nameRoom)

	err := false
	errMessage := ""

	if !common.Contain(typeRoom,common.Commandes) {
		err = true
		errMessage = fmt.Sprintf("\nError ❌ : \nThe room (%v) invalid, unknown command !",typeRoom)
	} else if nameRoom == "" {
		err = true
		errMessage = ("\nError ❌ : \nThe room must have a name !")
	} else if typeRoom == common.Commandes[common.INDEX_MIDDLEROOM] && len(ants) > 0 {
		err = true
		errMessage = ("\nError ❌ : \nThe middle room should not contain ants when initialized!")
	} else if typeRoom == common.Commandes[common.INDEX_END] && len(ants) > 0 {
		err = true
		errMessage = ("\nError ❌ : \nThe end room should not contain ants when initialized !")
	} else if typeRoom == common.Commandes[common.INDEX_START] && len(ants) == 0 {
		err = true
		errMessage = ("\nError ❌: \nThe start room must contain at least one ant !")
	}

	if err {
		return myRoom,errors.New(common.ColorString(common.INDEX_C_RED,errMessage))
	} else {
		myRoom.Type = typeRoom
		myRoom.Name = nameRoom
		myRoom.Point_X = point_x
		myRoom.Point_y = point_y
		myRoom.Ants = ants
	}

	fmt.Printf(common.ColorString(common.INDEX_C_GREEN,fmt.Sprintf(" ✅ The room ( %v - %v ) was successfully created !\n",typeRoom,nameRoom)))
	return myRoom,nil
}