package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 	fmt.Println("Defina o preço unitário da maçã:")

	// 	var precoUnitarioMaca string
	// 	//precoUnitarioMaca := ""
	// 	fmt.Scan(&precoUnitarioMaca)

	// 	fmt.Println(precoUnitarioMaca)

	// 	precounitarioMacaFloat, err := strconv.ParseFloat(precoUnitarioMaca, 64)
	// 	if err != nil {
	// 		fmt.Println("Erro ao converter preço da maçã de string para decimal", err.Error())
	// 		return
	// 	}

	// 	//fmt.Println(precounitarioMacaFloat)

	// 	fmt.Println("Quantos reais de maçã a cliente deseja?")

	// 	var pedidoCliente string
	// 	fmt.Scan(&pedidoCliente) //R$10.5

	// 	valorPagamento := strings.TrimPrefix(pedidoCliente, "R$")

	// 	valorPgtoFloat, err := strconv.ParseFloat(valorPagamento, 64)
	// 	if err != nil {
	// 		fmt.Println("Erro ao converter o valor do pagamento do cliente de string para decimal", err.Error())
	// 		return
	// 	}
	// 		qtdMacasAComprar := valorPgtoFloat / precounitarioMacaFloat
	// 		fmt.Println("Quantas maçãs é possível comprar?")

	// 		valorPgtoFormatado strconv.FormatFloat(valorPgtoFloat, 'f', 2, 64)

	// 		qtdMacasInt := int(qtdMacasAComprar)
	// 		fmt.Printf("A cliente pode comprar %d de maçãs com o valor de R$%s", qtdMacasInt, valorPgtoFormatado)
	// }
	// package main

	fmt.Println("Insira lista de números separados por vírgula sem espaço")

	var listaDeNumeros string
	var soma int
	var maior int

	fmt.Scan(&listaDeNumeros)
	listaDeNumerosSlice := strings.Split(listaDeNumeros, ",")
	fmt.Println(listaDeNumerosSlice)
	fmt.Println("Resultado:", listaDeNumerosSlice[0]+listaDeNumerosSlice[1])

	listaNumerosInt := make([]int, 0, len(listaDeNumeros))

	for _, numeroStr := range listaDeNumerosSlice {

		numeroInt, err := strconv.Atoi(numeroStr)

		if err != nil {
			fmt.Printf("erro ao fazer conversão de %s: %s\n", numeroStr, err.Error())

			//fmt.Println("pulando para o próximo número")

			continue
			//break
			//return
		}

		if maior < numeroInt {
			maior = numeroInt
		}

		soma += numeroInt

		listaNumerosInt = append(listaNumerosInt, numeroInt)
	}
	fmt.Println(listaNumerosInt)
	fmt.Println(soma)
	fmt.Println("Maior número da lista:", maior)
}
