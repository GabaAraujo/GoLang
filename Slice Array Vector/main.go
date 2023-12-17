package main

import (
	"fmt"
)

func main() {

	salFuncio := make(map[string]int) //voce atribui nome aos valores do vetor
	salFuncio["Gabriel"] = 69000
	salFuncio["Irineu"] = 10000

	println(salFuncio["Gabriel"])

	sal_val, existe := salFuncio["Irineu"] //retorna o valor e se ele existe
	fmt.Println(sal_val, existe)

	//salario := []int{1000,2000,3000,4000} //slide
	//salario := make([]int,5) //slice com valor fixo

	salario := []int{} //nao tem valor declarado

	/*
		for i := 0; i < 10; i++ {

			salario[i] = 100 + i //voce vai adicionando sem ter um valor fixo
		}
	*/ //funciona apenas em valores fixos

	for i := 0; i <= 10; i++ {

		salario = append(salario, 100+i) //voce vai adicionando sem ter um valor fixo, como se fosse uma atribuicao comum com variavel
	}

	for _, salario := range salario { //como toda variavel tem de ser utilizada em go
		fmt.Println(salario) //coloca o underline pra nao ter que ficar declarando, ele ja entende o tamanho da variavel
	}

	fmt.Println("Hello World")
}
