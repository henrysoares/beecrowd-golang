package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	number_of_sentences, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	for i := 0; i < int(number_of_sentences); i++ {
		scanner.Scan()
		current_sentence := scanner.Text()
		criptograph(current_sentence)
	}
}

func criptograph(raw_sentence string) {
	do_first_batch(&raw_sentence)
	do_second_batch(&raw_sentence)
	do_third_batch(&raw_sentence)
	fmt.Println(raw_sentence)
}

func do_first_batch(sentence *string) {
	var new_sentence string
	current_sentence := *sentence
	for i := 0; i < len(*sentence); i++ {
		ascii_letter := current_sentence[i]
		current_letter := string(current_sentence[i])
		new_letter := string(ascii_letter + 3)

		if (ascii_letter >= 65 && ascii_letter <= 90) || (ascii_letter >= 97 && ascii_letter <= 122) {
			new_sentence = new_sentence + new_letter
		} else {
			new_sentence = new_sentence + current_letter
		}
	}

	*sentence = new_sentence
}

func do_second_batch(sentence *string) {
	var inverted_string string
	for i := len(*sentence); i > 0; i-- {
		current_sentence := *sentence
		inverted_string = inverted_string + string(current_sentence[i-1])
	}

	*sentence = inverted_string
}

func do_third_batch(sentence *string) {
	half := len(*sentence) / 2.0
	size := len(*sentence)

	current_sentence := *sentence

	first_half := current_sentence[0:half]
	secound_half := current_sentence[half:size]

	var new_half string
	for i := 0; i < len(secound_half); i++ {
		ascii_letter := secound_half[i]
		next_letter := string(ascii_letter - 1)

		new_half = new_half + next_letter
	}

	*sentence = first_half + new_half
}
