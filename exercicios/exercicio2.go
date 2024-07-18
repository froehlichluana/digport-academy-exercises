package exercicios

import (
	"fmt"
)

type Contato struct{
	Nome string
	Telefone string
	Email string
}

func ListaDeContatos() {

	contato1 := Contato{Nome: "Lila", Telefone: "274029865", Email: "lila@example.com"}
	contato2 := Contato{Nome: "Lenu", Telefone: "7642909362", Email: "lenu@example.com"}

	contatos := []Contato{contato1, contato2}

	for i, contato := range contatos {
		fmt.Printf("Contato %d:\n", i+1)
		fmt.Printf("Nome: %s\nTelefone: %s\nEmail: %s\n\n", contato.Nome, contato.Telefone, contato.Email)
	}

	contato3 := Contato{Nome: "Nino", Telefone: "537402830", Email: "nino@example.com"}
	contatos = append(contatos, contato3)
}