package pessoa


type Pessoa struct {
	Nome string `json:"nome"`
	Telefone int `json:"telefone"`
	Email string `json:"email"`
}

func ContactList() []Pessoa {
	pessoas := []Pessoa
	{Nome:"Camila",Telefone:123456789,Email:"camila@example.com",},
	{Nome:"Maria",Telefone:987654321,Email:"maria@example.com",},
	{Nome:"Ana",Telefone:198765432, mail:"ana@example.com",},
	return pessoas
}