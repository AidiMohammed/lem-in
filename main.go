package main

import (
	"lem-in/tools"
	"fmt"
)

func main(){
	anthill,err := tools.LaodFileInput("./mohammed.txt")
	tools.HandelError(err,"")

	err = anthill.ValidateAnthill()
	
	if err != nil {
		fmt.Println(err)
	} else {
		anthill.ShowAnthill()
	}
}