package main

import "fmt"

type Pessoa struct {
	nome string
}

	func (p Pessoa) digaAlo() {
		fmt.Println("Alo", p.nome)
	}

	func main(){
		p:= Pessoa{nome:"Alice"}
		p2 := Pessoa{nome: "Maria"}

		p.digaAlo()
		p2.DigaAlo()
	}
}