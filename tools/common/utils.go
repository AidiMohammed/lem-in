package common

func Contain(str string, strs []string) bool {
	for _, commande := range strs {
		if commande == str {
			return true
		}
	}
	return false
}