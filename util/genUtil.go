package util

import "strconv"

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
