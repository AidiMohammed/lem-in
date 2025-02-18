package main

import (
	"lem-in/tools"
	"fmt"
)

func main(){
	myAnthill,err := tools.LaodFileInput("./input01.txt")
	tools.HandelError(err,"")
	
	err = myAnthill.ValidateAnthill()

	if err != nil {
		fmt.Println(err)
	} else {
		myAnthill.ShowAnthill()
	}
}