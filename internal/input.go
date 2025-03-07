package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() PasswordOptions {
	reader := bufio.NewReader(os.Stdin)

	// Solicita o tamanho da senha
	fmt.Print("Digite o tamanho da senha: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	length, err := strconv.Atoi(lengthInput)
	if err != nil || length <= 0 {
		fmt.Println("Entrada inválida! Usando tamanho padrão: 8")
		length = 8
	}

	// Pergunta quais tipos de caracteres incluir
	options := PasswordOptions{Length: length}
	options.Uppercase = askUser("Incluir letras maiúsculas ? (s/n): ")
	options.Lowercase = askUser("Incluir letras minúsculas ? (s/n): ")
	options.Numbers = askUser("Incluir números ? (s/n): ")
	options.Specials = askUser("Incluir caracteres especiais ? (s/n): ")

	return options
}

func askUser(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(question)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if "s" == input {
			return true
		} else if "n" == input {
			return false
		} else {
			fmt.Println("Resposta inválida! Digite 's' para sim ou 'n' para não.")
		}
	}
}
