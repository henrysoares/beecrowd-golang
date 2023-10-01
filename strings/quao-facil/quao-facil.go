package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
*

TopCoder decidiu automatizar o processo de atribuição de níveis de dificuldade para os problemas. Os desenvolvedores do TopCoder concluíram que a dificuldade do problema esta relacionado apenas ao comprimento médio das palavras do enunciado do problema. Se o comprimento médio das palavras do enunciado é menor ou igual a 3, o problema recebe dificuldade de 250 pontos. Se o comprimento médio das palavras do enunciado for 4 ou 5, o problema recebe dificuldade de 500 pontos. Se o comprimento médio das palavras do enunciado for maior ou igual a 6, o problema recebe dificuldade de 1000 pontos.

Definições:

Símbolo: um conjunto de carateres ligados em ambos os lados por espaços, ou pelo início da descrição do problema, ou ainda pelo fim da descrição do problema.

Palavra: um símbolo que contenha apenas letras a-z ou A-Z, e pode terminar com um único ponto.

Comprimento da palavra: número de letras de uma palavra (um ponto não é uma letra).

Exemplos de símbolos que são palavras (aspas duplas apenas para exemplificar): "AB", "ab".

Exemplo de símbolos que não são palavras: "ab..", "a.b", ".ab", "a.b.", "a2b.", ".".

O comprimento médio das palavras é dado pela soma dos tamanhos das palavras do enunciado dividido pelo numero de palavras, a divisão é feita por números inteiros. Se o número de palavras for zero, então o comprimento médio das palavras é zero.

Sua tarefa é dado o enunciado do problema, computar a sua classificação de dificuldade do problema, que poderá ser 250, 500, ou 1000.
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(calculate_difficulty(input))
	}
}

func calculate_difficulty(input string) int {
	words := strings.Fields(input)
	var valid_words int
	var words_total_lenght int

	for _, word := range words {
		if is_valid_word(word) {
			valid_words++
			words_total_lenght += len(strings.ReplaceAll(word, ".", ""))
		}
	}

	if valid_words > 0 {
		average := words_total_lenght / valid_words

		switch {
		case average <= 3:
			return 250
		case average <= 5:
			return 500
		default:
			return 1000
		}
	}

	return 250
}

func is_valid_word(word string) bool {
	for i, c := range word {
		ascii_letter := int(c)
		is_on_lower_ascii_range := (ascii_letter >= 97 && ascii_letter <= 122)
		is_on_upper_ascii_range := (ascii_letter >= 65 && ascii_letter <= 90)
		is_final_period := string(c) == "." && i == len(word)-1

		if !(is_on_lower_ascii_range || is_on_upper_ascii_range || is_final_period) {
			return false
		}
	}

	return true
}
