package common

import (
	"fmt"
	"errors"
)

const INDEX_START = 0
const INDEX_END = 1
const INDEX_COMMENT = 2
const INDEX_MIDDLEROOM = 3

const INDEX_C_RED = 31
const INDEX_C_GREEN = 32
const INDEX_C_YELLOW = 33
const INDEX_C_BLUE = 34

var Commandes = []string{"##start","##end","##comment","middleroom"}

func Contain(str string, strs []string) bool {
	for _, str := range strs {
		if str == str {
			return true
		}
	}
	return false
}

func ColorString (indexColor int,str string) string{
	return fmt.Sprintf("\033[%vm%v\033[0m",indexColor,str)
}

func BoxString(str string) error {

	if len(str) > 0 {
		str = "==== "+str+" ===="
		for index := 0; index < 2 ;index++ {
			for indexStr := 1;indexStr < len(str);indexStr++ {
				fmt.Print("=")
			}
			if(index == 0){
				fmt.Println("\n"+str)
			} else {
				fmt.Println()
			}
		}
		return nil
	}

	return errors.New("The input string is empty")
}