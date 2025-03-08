package main

import (
	//"fmt"
	"lem-in/tools"
	"lem-in/tools/common"
	//"lem-in/room"
	"lem-in/anthill"
)

func main() {
	var myAnthill anthill.Anthill

	myAnthill,err := tools.MakeAnthill("input01.txt")
	errMessage := common.ColorString(common.INDEX_C_RED,"\nFailed to create the anthill, Please check your input file end try again later")
	if tools.HandelError(err,errMessage){
		return
	} else {	
		myAnthill.ShowAnthill()
	}
} 