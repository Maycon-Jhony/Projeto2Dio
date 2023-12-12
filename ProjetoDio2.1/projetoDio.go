package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		output := ""

		// Verifica se é múltiplo de 3
		if i%3 == 0 {
			output += "Pin"
		}

		// Verifica se é múltiplo de 5
		if i%5 == 0 {
			output += "Pan"
		}

		// Imprime o número ou a combinação
		if output == "" {
			fmt.Println(i)
		} else {
			fmt.Println(output)
		}
	}
}
