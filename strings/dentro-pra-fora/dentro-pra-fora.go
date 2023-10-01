package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
*
A sua impressora foi infectada por um vírus e está imprimindo de forma incorreta. Depois de olhar para várias páginas impressas por um tempo, você percebe que ele está imprimindo cada linha de dentro para fora. Em outras palavras, a metade esquerda de cada linha está sendo impressa a partir do meio da página até a margem esquerda. Do mesmo modo, a metade direita de cada linha está sendo impressa à partir da margem direita e prosseguindo em direção ao centro da página.

Por exemplo a linha:
THIS LINE IS GIBBERISH

está sendo impressa como:
I ENIL SIHTHSIREBBIG S

Da mesma foma, a linha " MANGOS " está sendo impressa incorretamente como "NAM  SOG". Sua tarefa é desembaralhar (decifrar) a string a partir da forma como ela foi impressa para a sua forma original. Você pode assumir que cada linha conterá um número par de caracteres.
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	cases, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	for i := 0; int64(i) < cases; i++ {
		scanner.Scan()

		sentence := scanner.Text()

		first_half := sentence[0 : len(sentence)/2]
		second_half := sentence[len(sentence)/2 : len(sentence)]

		fmt.Println(invert_string(first_half) + invert_string(second_half))
	}
}

func invert_string(str string) string {
	var new_str string
	for i := len(str) - 1; i >= 0; i-- {
		current_value := string(str[i])
		new_str = new_str + current_value
	}

	return new_str
}
