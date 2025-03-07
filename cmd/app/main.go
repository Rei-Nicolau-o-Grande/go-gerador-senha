package main

import (
	"fmt"
	"github.com/Rei-Nicolau-o-Grande/generator-password/internal"
)

func main() {
	options := internal.Input()

	password, err := internal.GeneratePassword(options)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n🔑 Senha gerada:", password)
}
