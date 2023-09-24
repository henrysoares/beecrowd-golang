package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
*
Questão:

Solicitaram para que você construisse um programa simples de criptografia. Este programa deve possibilitar enviar mensagens codificadas sem que alguém consiga lê-las. O processo é muito simples. São feitas três passadas em todo o texto.

Na primeira passada, somente caracteres que sejam letras minúsculas e maiúsculas devem ser deslocadas 3 posições para a direita, segundo a tabela ASCII: letra 'a' deve virar letra 'd', letra 'y' deve virar caractere '|' e assim sucessivamente. Na segunda passada, a linha deverá ser invertida. Na terceira e última passada, todo e qualquer caractere a partir da metade em diante (truncada) devem ser deslocados uma posição para a esquerda na tabela ASCII. Neste caso, 'b' vira 'a' e 'a' vira '`'.

Por exemplo, se a entrada for “Texto #3”, o primeiro processamento sobre esta entrada deverá produzir “Wh{wr #3”. O resultado do segundo processamento inverte os caracteres e produz “3# rw{hW”. Por último, com o deslocamento dos caracteres da metade em diante, o resultado final deve ser “3# rvzgV”.
Entrada

A entrada contém vários casos de teste. A primeira linha de cada caso de teste contém um inteiro N (1 ≤ N ≤ 1*104), indicando a quantidade de linhas que o problema deve tratar. As N linhas contém cada uma delas M (1 ≤ M ≤ 1*103) caracteres.
Saída

Para cada entrada, deve-se apresentar a mensagem criptografada.
*/
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
