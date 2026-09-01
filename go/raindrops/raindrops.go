package raindrops

import "fmt"

func Convert(number int) string {
	palavra := ""
	if number%3 == 0 {
		palavra += "Pling"
	}
	if number%5 == 0 {
		palavra += "Plang"
	}
	if number%7 == 0 {
		palavra += "Plong"
	}
	if palavra == "" {
		palavra2 := fmt.Sprintf("%d", number)
		return palavra2
	}

	return palavra
}
