package logic

import "fmt"

func Tabuada() {
	var numerador int 
	var resultado int
	fmt.Println("Qual número você quer calcular a tabuada?")
	fmt.Scanf("%d", &numerador)

	

	for i := 1; i <= 10; i++ {
		
		resultado = numerador * i
		fmt.Printf("%d X %d = %d/n", i, numerador, resultado)
	}
}