package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var digit, number int
		fmt.Scan(&digit, &number)

		if digit == 0 && number == 0 {
			break
		}

		new_value := corrigirValor(digit, number)
		fmt.Println(new_value)
	}
}

func corrigirValor(digit int, number int) int {
	formatted_number := strconv.Itoa(number)
	str_value := ""

	for _, current_value := range formatted_number {
		d, _ := strconv.Atoi(string(current_value))
		if d != digit {
			str_value += string(current_value)
		}
	}

	if str_value == "" {
		return 0
	}

	final_response, _ := strconv.Atoi(str_value)
	return final_response
}
