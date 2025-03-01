package tools

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"errors"
	"lem-in/anthill"
	"lem-in/room"
	"lem-in/tools/common"

)

func MakeAnthill(filePath string)(anthill.Anthill,error){
	file, err := os.Open(filePath)
	HandelError(err,"")

	isFlag := false
	var theFlag string
	var rooms []room.Room
	var myAnthill anthill.Anthill
	var indexLine uint
	var ants []uint
	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		indexLine++

		if indexLine == 1 {
			intAnts,err := strconv.Atoi(line)

			for index := 0 ; index < intAnts ; index++ {
				ants = append(ants,uint(index))
			}
			if HandelError(err,"") {
				return myAnthill,errors.New(common.ColorString(common.INDEX_C_RED,"Error ❌ : \nYour file dose not contain a number of ants !"))
			}
			continue
		}

		if strings.HasPrefix(line,"##") {
			if line == common.Commandes[common.INDEX_START] || line == common.Commandes[common.INDEX_END] {
				isFlag = true
				theFlag = line
				continue
			} else if line == common.Commandes[common.INDEX_COMMENT] {
				continue
			} else {
				errMessage := fmt.Sprintf("Error ❌ : \nInvalid input line : %v (%v)",indexLine,line)
				return myAnthill,errors.New(common.ColorString(common.INDEX_C_RED,errMessage))
			}
		}

		lineSplit := strings.Split(line," ")
		if len(lineSplit) > 3 {
			return myAnthill,errors.New(common.ColorString(common.INDEX_C_RED,fmt.Sprintf("Error ❌ : \nInvalid format file line %v : %v",indexLine,line)))
		} else if len(lineSplit) == 1 {
			lineSplit = strings.Split(line,"-")
			if len(lineSplit) == 2 {
				//Tunnel handling

				continue
			} else {
				return myAnthill,errors.New(common.ColorString(common.INDEX_C_RED,fmt.Sprintf("Error ❌ : \nInvalid format file line %v : %v",indexLine,line)))
			}
		}

		if isFlag {
			//Flag Rooms (start/end) handling
			nameRoom := lineSplit[0] 
			antsRoom := []uint{}
			if theFlag == common.Commandes[common.INDEX_START] {
				antsRoom = ants
				ants = []uint{}
			} 

			point_x,err := strconv.Atoi(lineSplit[1])
			if err != nil{
				return myAnthill,errors.New(fmt.Sprintf("Error ❌ : \n%v",err))
			}
			point_y,err := strconv.Atoi(lineSplit[2])
			if err != nil{
				return myAnthill,errors.New(fmt.Sprintf("Error ❌ : \n%v",err))
			}
			myRoom,err := room.MakeRoom(theFlag, nameRoom, point_x, point_y, antsRoom)
			if HandelError(err,"") {
				return myAnthill,err
			}
			rooms = append(rooms,myRoom)
			isFlag = false
			theFlag = ""
			continue
		} else {
			//Middle rooms handling
			nameRoom := lineSplit[0]

			point_x,err := strconv.Atoi(lineSplit[1])
			if err != nil{
				return myAnthill,errors.New(fmt.Sprintf("Error ❌ : \n%v",err))
			}
			point_y,err := strconv.Atoi(lineSplit[2])
			if err != nil{
				return myAnthill,errors.New(fmt.Sprintf("Error ❌ : \n%v",err))
			}
			myRoom,err := room.MakeRoom(common.Commandes[common.INDEX_MIDDLEROOM], nameRoom, point_x, point_y, []uint{})
			if HandelError(err,"") {
				return myAnthill,err
			}
			rooms = append(rooms,myRoom)
		}
	}

	err = myAnthill.InitAnthill(rooms)
	if HandelError(err,""){
		return myAnthill,errors.New("")
	}
	return myAnthill,nil
}

func HandelError(err error,customMessage string) bool{	
	if err != nil {
		if customMessage != "" {
			fmt.Printf("%v",customMessage)
			return true
		} else {
			fmt.Printf("%v",err)
			return true
		}
	}
	return false
}

func getRoomByName(roomName string,rooms []room.Room) (*room.Room,error){
	for _,room := range rooms {
		if room.Name == roomName {
			return &room,nil
		}
	}
	emptyRoom := room.Room{}
	errMessage := fmt.Sprintf("Error ❌ : \nThe room %v is not found in the anthill ",roomName)
	return &emptyRoom,errors.New(common.ColorString(common.INDEX_C_RED,errMessage)) 
}