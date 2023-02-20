package util

import (
	"strconv"
)

func GenToNumeral(genNumber interface{}) string {
	genNumberInt := 9
	switch value := genNumber.(type) {
	case string:
		genNumberConv, err := strconv.Atoi(value)
		if err != nil {
			return "Latest"
		}
		genNumberInt = genNumberConv
	case int:
		genNumberInt = value
	}

	if genNumberInt == 4 {
		return "IV"
	} else if genNumberInt < 1 || genNumberInt > 9 {
		return "Latest"
	}

	numeral := ""
	if genNumberInt >= 5 {
		numeral += "V"
		genNumberInt -= 5
	}
	for i := 0; i < 3; i++ {
		if genNumberInt > 0 {
			numeral += "I"
			genNumberInt--
		}
	}
	return numeral
}
