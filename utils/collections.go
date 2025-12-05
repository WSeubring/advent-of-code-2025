package utils

func RemoveIndexFromArr[Type any](s []Type, index int) []Type {
	left := append([]Type{}, s[:index]...)
	right := append([]Type{}, s[index+1:]...)
	return append(left, right...)
}
