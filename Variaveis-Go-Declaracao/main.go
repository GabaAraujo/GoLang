package main

import (
	"fmt"
)

func main() {

	name := getName()
	fmt.Println("Hello", name)

	//var salario, bonus float32 = 100,salario * 0.20 nao pode declaracao dupla usando a variavel na segunda

	var salario float32 = 100
	var bonus float32 = salario * 0.20
	novo_salario, new_bonus := newSalario(salario, bonus) //: serve para receber valores

	fmt.Println("novo salario:", novo_salario, "novo bonus:", new_bonus)

}

func getName() string { //retorna uma string

	return "Gabriel"
}

func newSalario(ValorAtual float32, bonus float32) (float32, float32) {

	return ValorAtual + bonus, bonus

}
