package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
*
Encontre a maior substring comum entre as duas strings informadas. A substring pode ser qualquer parte da string, inclusive ela toda. Se não houver subseqüência comum, a saída deve ser “0”. A comparação é case sensitive ('x' != 'X').
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		first_word := "abcdef"
		scanner.Scan()
		second_word := "cdofhij"
		maximum := 0

		if len(second_word) > len(first_word) {
			biggest_word := second_word
			second_word = first_word
			first_word = biggest_word
		}

		for i := 0; i < len(second_word); i++ {
			for j := 0; j < i+1; j++ {
				current_sentence := second_word[j : len(second_word)-i+j]

				if strings.Contains(first_word, current_sentence) {
					maximum = max(maximum, len(second_word)-i)
				}
			}
		}

		fmt.Println(maximum)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
