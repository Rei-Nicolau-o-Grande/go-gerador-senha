package internal

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PasswordOptions struct {
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
	Specials  bool
}

const (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "0123456789"
	specials  = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

func GeneratePassword(options PasswordOptions) (string, error) {
	var charset string

	if options.Length == 0 {
		options.Length = 8
	}

	if options.Uppercase {
		charset += uppercase
	}

	if options.Lowercase {
		charset += lowercase
	}

	if options.Numbers {
		charset += numbers
	}
	if options.Specials {
		charset += specials
	}

	if len(charset) == 0 {
		return "", fmt.Errorf("nenhuma opção de caractere foi selecionada")
	}

	password := make([]byte, options.Length)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}
