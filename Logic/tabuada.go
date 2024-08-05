package logic

import "fmt"

func Tabuada() {
	var numerador int 
	var resultado int

	fmt.Println("Qual número você quer calcular a tabuada?")
	fmt.Scan(numerador)

	for i := 1; i <= 10; i++{
		resultado = numerador * i
		fmt.Print(i, "X", numerador, " = ", resultado)
	}
}