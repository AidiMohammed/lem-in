package tools

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"errors"
)
var flags = []string{"##start","##end","##comment"}
type Anthill struct {
	StartRoom []int
	EndRoom []int
	isMiddelRooms [][]int
}

func (this Anthill) ValidateAnthill() error {
	var err error
	if len(this.StartRoom) == 0 || len(this.EndRoom) == 0 {
		err = errors.New("The Room ##start or ##end is required")
	} else if len(this.isMiddelRooms) == 0 {
		err = errors.New("you must have at least one middle room")
	}

	if err != nil {
		return errors.New(fmt.Sprintf("The Anthill is not valid : %v",err))
	} else {
		return nil
	}
}

func (this Anthill) ShowAnthill() {

	fmt.Println("====================================")
	fmt.Println("  📜 Information  de la Foumilière  ")
	fmt.Println("====================================")

	fmt.Printf("Start Room : %v \n",this.StartRoom)
	fmt.Printf("Middle Rooms : %v\n",this.isMiddelRooms)
	fmt.Printf("End Room : %v \n",this.EndRoom)
	
	fmt.Println("====================================")
}

func LaodFileInput(filePath string)(Anthill,error){
	isFlag := false
	theFlag := ""
	isMiddelRoom := false
	myAnthill := Anthill{}
	indexLine := 0

	file, err := os.Open(filePath)
	HandelError(err,"")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		indexLine++
		line := scanner.Text()

		if isFlag {
			if theFlag == flags[0] || theFlag == flags[1] {
				target :=  strings.Split(line," ")
				if len(target) != 3 {
					return myAnthill,errors.New(fmt.Sprintf("Invalid input Line : %v ( %v )\n",indexLine,line))
				}
				for _, strVal := range target {
					intVal, err := strconv.Atoi(strVal)
					if err != nil {
						return myAnthill,err
					}
					if theFlag == flags[0] {
						myAnthill.StartRoom = append(myAnthill.StartRoom, intVal)
					} else {
						myAnthill.EndRoom = append(myAnthill.EndRoom, intVal)
					}
					if theFlag == flags[0] { //Start Room
						isMiddelRoom = true
					} else if theFlag == flags[1] { //End Room
						isMiddelRoom = false
					}
				}
				isFlag = false
				theFlag = ""
				continue
			}
		}

		if strings.HasPrefix(line,"##") {
			if line == flags[0] || line == flags[1] {
				theFlag = line
				isFlag = true
				continue			
			} else if line == flags[2] {
				continue
			} else {
				return myAnthill,errors.New(fmt.Sprintf("Invalid input Line : %v ( %v )\n",indexLine,line))
			}
		}

		if isMiddelRoom {
			target := strings.Split(line," ")
			if len(target) != 3 {
				return myAnthill,errors.New(fmt.Sprintf("Invalid input Line : %v ( %v )\n",indexLine,line))
			}
			var newMiddleRoom []int
			for _,strVal := range target {
				intVal,err := strconv.Atoi(strVal) 
				if err != nil {
					return myAnthill,err
				}
				newMiddleRoom = append(newMiddleRoom,intVal)
			}
			myAnthill.isMiddelRooms = append(myAnthill.isMiddelRooms,newMiddleRoom)
		}
	}
	return myAnthill,nil
}

func HandelError(err error,message string) {	
	if err != nil {
		if message != "" {
			fmt.Printf("%v",message)
		} else {
			fmt.Printf("%v",err)
		}
	}
}

