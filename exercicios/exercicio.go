package exercicios 

import (
	"fmt"
)


func Despesas() {
minhasDespesas := []string{"mercado", "aluguel", "condominio", "eletricidade", "combustivel", "internet"} 

for _, despesa := range minhasDespesas{
	fmt.Println(despesa)
}

fmt.Println("Qual item você deseja saber se está na sua lista de despesas?")

var itemVerificar string
fmt.Scan(&itemVerificar)

fmt.Println("Verificando se há", itemVerificar, "na sua lista...")
for _, v := range minhasDespesas{
	if v == itemVerificar{
		fmt.Println("Sim")
		return
	}
}
fmt.Println("Não")

fmt.Printf("Minha lista de despesas tem %d itens", len(minhasDespesas))

}