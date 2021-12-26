package util

import (
	"strconv"
	"strings"
)

func GenToNumber(genNumeral string) int {
	number := 0
	if strings.Contains(genNumeral, "V") {
		number += 5
	}

	number += strings.Count(genNumeral, "I")

	return number
}

func GenToNumeral(genNumber string) string {
	genNumberInt, err := strconv.Atoi(genNumber)
	if err != nil {
		return err.Error()
	}

	if genNumber == "4" {
		return "IV"
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
