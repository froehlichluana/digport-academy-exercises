package logic

import "fmt"

func Tabuada() {
	var numerador int 
	fmt.Println("Qual número você quer calcular a tabuada?")
	fmt.Scanf("%d", &numerador)
	i := 1

	for {
		if (i>10){
			break;
		}
		fmt.Println(numerador, " X ", i ," = ", numerador * i)
		i++
	}
}