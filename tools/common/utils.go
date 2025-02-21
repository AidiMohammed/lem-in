package common

import "fmt"

const INDEX_C_RED = 31
const INDEX_C_GREEN = 32
const INDEX_C_YELLOW = 33
const INDEX_C_BLUE = 34

func Contain(str string, strs []string) bool {
	for _, commande := range strs {
		if commande == str {
			return true
		}
	}
	return false
}

func ColorString (indexColor int,str string) string{
	return fmt.Sprintf("\033[%vm%v\033[0m",indexColor,str)
}