package tools

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"errors"
	"lem-in/anthill"
)

func LaodFileInput(filePath string)(anthill.Anthill,error){
	isFlag := false
	theFlag := ""
	myAnthill := anthill.Anthill{}
	myRoom := myAnthill.Rooms
	indexLine := 0

	file, err := os.Open(filePath)
	HandelError(err,"")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		indexLine++
		line := scanner.Text()

		if indexLine == 1 {
			ants,err := strconv.Atoi(line)
			if err != nil {
				myAnthill = anthill.Anthill{}
				return myAnthill,err
			}
			myAnthill.AddAntsInRoom(ants,anthill.Flags[0],"")
			myRoom = myAnthill.GetRoomType(myAnthill)
			continue
		}

		if strings.HasPrefix(line,"##") {
			if line == anthill.Flags[0] || line == anthill.Flags[1] {
				theFlag = line
				isFlag = true
				continue
			} else if line == anthill.Flags[2] {
				continue
			} else {
				myAnthill = anthill.Anthill{}
				return myAnthill,errors.New(fmt.Sprintf("Invalid input Line : %v ( %v )\n",indexLine,line))
			}
		}

		if isFlag {
			if theFlag == anthill.Flags[0] || theFlag == anthill.Flags[1] {
				err := makeRoom(&myRoom,theFlag,line, indexLine)
				if !HandelError(err,"") {
					myAnthill.AddRoomInAnthill(myRoom)
					isFlag = false
					theFlag = ""
					myRoom = anthill.Room{}
					continue
				} else {
					myAnthill = anthill.Anthill{}
					return myAnthill,err
				}

			}
		}

		if !isFlag {
			err := makeRoom(&myRoom,"middleRoom",line, indexLine)
			if !HandelError(err, "") {
				myAnthill.AddRoomInAnthill(myRoom)
				isFlag = false
				theFlag = ""
				//myRoom = anthill.Room{}
			} else {
				myAnthill = anthill.Anthill{}
				return myAnthill,err
			}
		}

	}
	return myAnthill,nil
}

func HandelError(err error,customMessage string) bool{	
	if err != nil {
		if customMessage != "" {
			fmt.Printf("%v",customMessage)
		} else {
			fmt.Printf("%v",err)
		}
		return true
	}
	return false
}


