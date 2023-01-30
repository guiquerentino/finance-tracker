package main

import (
	"financial-tracker/structs"
	"financial-tracker/utils"
	"fmt"
	"os"
	"reflect"
	"time"
)

var (
	bd        []structs.Transaction
	selection int
)

func main() {
	MainMenu()
}

func MainMenu() {

	fmt.Println("O que deseja fazer?\n 1 - Adicionar transação \n 2 - Ver transações \n 3 - Remover transação")

	fmt.Scanln(&selection)

	switch selection {
	case 1:
		AddTransaction()
	case 2:
		AllTransactions()
	case 3:
		DeleteTransaction()
	}

}

func AddTransaction() {

	var transactionName string
	var transactionValue float64

	fmt.Print("Digite o nome da transação:")
	fmt.Scanln(&transactionName)

	if reflect.TypeOf(transactionName) != reflect.TypeOf(string("a")) {
		fmt.Println("Valor inválido!")
		AddTransaction()
	}

	fmt.Print("Digite o valor da transação:")
	fmt.Scanln(&transactionValue)

	if reflect.TypeOf(&transactionValue) != reflect.TypeOf(float64(0)) || reflect.TypeOf(&transactionName) != reflect.TypeOf(int(0)) {
		fmt.Println("Valor inválido!")
		AddTransaction()
	}

	t := structs.NewTransaction(transactionValue, transactionName)
	bd = append(bd, *t)

	fmt.Println("Transferencia criada com sucesso! ")

	fmt.Println("Deseja voltar ao menu? \n 1 - Sim\n 2 - Não")
	fmt.Scanln(&selection)

	if selection == 1 {
		utils.ClearScreen()
		MainMenu()
	} else if selection == 2 {
		os.Exit(1)
	}

	utils.ClearScreen()
	MainMenu()

}

func AllTransactions() {
	total := 0.0

	for _, j := range bd {
		fmt.Println("---------------------------------------------------")
		fmt.Println("Nome da transação -> ", j.GetTransactionName())
		fmt.Printf("Valor da transação -> %g\n", j.GetValue())
		fmt.Println("Data da transação -> ", j.GetTransactionDate().Format(time.RFC850))
		fmt.Println("ID da Transação -> ", j.GetTransactionId())
		fmt.Println("---------------------------------------------------")
		total += j.GetValue()
	}
	fmt.Printf("Total de despesa: %g\n", total)

	fmt.Println("Deseja voltar ao menu? \n 1 - Sim\n 2 - Não")
	fmt.Scanln(&selection)

	if selection == 1 {
		utils.ClearScreen()
		MainMenu()
	} else if selection == 2 {
		os.Exit(1)
	}

	utils.ClearScreen()
	MainMenu()

}

func DeleteTransaction() {
	var transactionName string

	fmt.Print("Digite o nome da transação a ser removida: ")
	fmt.Scanln(&transactionName)

	for i, j := range bd {
		if j.GetTransactionName() == transactionName {
			bd = append(bd[:i], bd[i+1:]...)
			break
		}
	}

	fmt.Println("Transação removida com sucesso!")

	utils.ClearScreen()
	MainMenu()
}
