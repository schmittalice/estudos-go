package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Livro struct {
	Nome    string `json:"nome"`
	Editora string `json:"editora"`
	Autor   string `json:"autor"`
	Link    string `json:"link"`
}

func main() {
	jsonFile, err := os.Open("livro.json")
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)
	bytesJson, err := io.ReadAll(jsonFile)

	livro := Livro{}

	err = json.Unmarshal(bytesJson, &livro)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(livro.Nome)

	livro.Nome = "Livro meninas poderosas nova edicao"
	fmt.Println(livro.Nome)
	fmt.Println(livro)

	newBytesJson, err := json.Marshal(livro)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("livro.json", newBytesJson, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Arquivo criado com sucesso")
}
