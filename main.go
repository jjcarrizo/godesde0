package main

import (
	"fmt"

	"github.com/jjcarrizo/godesde0/variables"
)

func main() {
	estado, texto := fmt.Println(variables.ConvertiraTexto(123))
	fmt.Println(estado)
	fmt.Println(texto)
}
