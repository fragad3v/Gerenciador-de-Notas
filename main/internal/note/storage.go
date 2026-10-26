package main

import (
	"os"
	"fmt"
)

func main(){
	nomeArquivo := "notas.txt"
	conteudoEscrita := "Ola! Gerando sua nota!"
	
	// 1. ESCREVENDO EM UM ARQUIVO
	err := os.WriteFile(nomeArquivo, []byte(conteudoEscrita),0644)
	if err != nil {
		fmt.Printf("Erro ao escrever o arquivo: %v", err)
	}

	fmt.Println("Arquivo gravado com sucesso!")

	// 2. LENDO DE UM ARQUIVO
	dadosByte, err := os.ReadFile(nomeArquivo)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v", err)
		return
	}
	fmt.Printf("Conteúdo lido: %s\n", string(dadosBytes))
}