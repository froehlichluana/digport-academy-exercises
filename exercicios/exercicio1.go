package exercicios 

import (
	"fmt"
)


var mercado float64 = 945.89
var aluguel float64 = 1200.00
var condominio float64 = 395.67
var eletricidade float64 = 145.78
var combustivel float64 = 543.90



func TotalDespesas() {
var minhasDespesas =[]float64{mercado, aluguel, condominio, eletricidade, combustivel} 
for _, v := range minhasDespesas{
	fmt.Println(v)
}

}