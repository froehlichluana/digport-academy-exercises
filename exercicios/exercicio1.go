package exercicios

import "fmt"

func TotalDepesas() {
	despesas := map[string]int{"mercado":1100, "aluguel":1500, "gás":90, "transporte":900, "internet":119}

	total := 0

	var get_despesa string
	fmt.Println("Digite a despesa: ")
	fmt.Scan(&get_despesa)

	v, encontrada := despesas[get_despesa]
	if encontrada {
		fmt.Printf("O valor da despesa %s informada é: %d\n", get_despesa, v)
	} else {
		fmt.Printf("despesa %s não encontrada \n", get_despesa)
	}

	for despesa, valor := range despesas {
		total = total + valor
		fmt.Printf("A despesa %s é %d\n", despesa, valor)
	}

	budget := 5500
	if total > budget{
		fmt.Printf("O valor total das despesas %d ultrapassou o seu orçamento \n", total)
	}
	fmt.Printf("O total das despesas é %d\n", total)


}