package utils

import (
	"fmt"
	"strconv"
)

func IsNumber(str string) int {
	matched := MatchString(`^[0-9]*$`, str)
	if !matched {
		return -1
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("[utils IsNumber err ] strconv.Atoi ", err.Error())
		return -1
	}
	return num
}
