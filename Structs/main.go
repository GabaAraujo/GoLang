package main

import (
	"fmt"
)

type Pessoa struct {
	Nome    string
	Idade   int
	Salario float32
}

func main() {

	pessoa := Pessoa{

		Nome:    "Gabriel",
		Idade:   25,
		Salario: 10000,
	}

	fmt.Println("Salario antigo:", pessoa.Salario)

	novo_Salario(&pessoa)

	//pessoa = novo_Salario(pessoa) //atribuindo a copia enviada na variavel

	fmt.Println("Salario Novo:", pessoa.Salario)

}

/* quando vc passa uma sctruct que nao e o ponteiro, voce manda uma copia
func novo_Salario(p Pessoa) {

	var bonus float32 = p.Salario * 0.20
	p.Salario += bonus // salario + bonus

}
*/

func novo_Salario(p *Pessoa) {

	var bonus float32 = p.Salario * 0.20
	p.Salario += bonus // salario + bonus

}
