package main

import (
	"fmt"

	"github.com/Brayan-Coba/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
