package utils

func Delete(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
