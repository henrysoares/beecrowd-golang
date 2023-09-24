package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	number_of_sentences, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	for i := 0; i < int(number_of_sentences); i++ {
		scanner.Scan()
		current_sentence := scanner.Text()
		find_hidding_sentences(current_sentence)
	}
}

func find_hidding_sentences(sentence string) {
	sentence_separator := " "

	words := strings.Split(sentence, sentence_separator)

	for _, word := range words {
		if word != "" {
			fmt.Printf("%c", word[0])
		}
	}

	fmt.Println("")
}
