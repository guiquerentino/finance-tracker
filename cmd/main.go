package main

import (
	"financial-tracker/structs"
	"financial-tracker/utils"
	"fmt"
	"gorm.io/gorm"
	"os"
	"reflect"
	"time"
)

var selection int

func main() {

	db, err := utils.ConnectDB()

	if err != nil {
		panic(err)
	}

	MainMenu(db)
}

func MainMenu(db *gorm.DB) {

	fmt.Println("Select an option?\n 1 - Add transaction \n 2 - Check transactions \n 3 - Remove transaction\n 4 - Update transaction")

	_, err := fmt.Scanln(&selection)

	if err != nil {
		utils.ClearScreen()
		MainMenu(db)
	}

	switch selection {
	case 1:
		AddTransaction(db)
	case 2:
		AllTransactions(db)
	case 3:
		DeleteTransaction(db)
	case 4:
		UpdateTransaction(db)
	default:
		utils.ClearScreen()
		MainMenu(db)
	}

}

func AddTransaction(db *gorm.DB) {

	var transactionName string
	var transactionValue float64

	fmt.Print("Type the transaction name:")

	_, err := fmt.Scanln(&transactionName)
	if err != nil {
		AddTransaction(db)
	}

	if reflect.TypeOf(transactionName) != reflect.TypeOf("a") {
		fmt.Println("Invalid value!")
		AddTransaction(db)
	}

	fmt.Print("Type the transaction value:")

	_, err = fmt.Scanln(&transactionValue)

	if err != nil {
		AddTransaction(db)
	}

	if reflect.TypeOf(transactionValue) != reflect.TypeOf(float64(0)) {
		fmt.Println("Invalid value!!")
		AddTransaction(db)
	}

	t := structs.NewTransaction(transactionValue, transactionName)

	db.Create(&t)

	fmt.Println("Transaction successfully created! ")

	fmt.Println("Go back to the menu? \n 1 - Yes\n 2 - No")

	_, err = fmt.Scanln(&selection)

	if err != nil {
		utils.ClearScreen()
		MainMenu(db)
	}

	if selection == 1 {
		utils.ClearScreen()
		MainMenu(db)
	} else if selection == 2 {
		os.Exit(1)
	}

	utils.ClearScreen()
	MainMenu(db)

}

func AllTransactions(db *gorm.DB) {
	total := 0.0

	var objetos []structs.Transaction

	db.Find(&objetos)

	for _, j := range objetos {
		fmt.Println("---------------------------------------------------")
		fmt.Println("Transaction name -> ", j.TransactionName)
		fmt.Printf("Transaction value -> %g\n", j.Value)
		fmt.Println("Transaction Date -> ", j.TransactionDate.Format(time.RFC822))
		fmt.Println("Transaction ID -> ", j.ID)
		total += j.Value
	}

	fmt.Printf("Total expenses: %g\n", total)

	fmt.Println("Go back to the menu? \n 1 - Yes\n 2 - No")

	_, err := fmt.Scanln(&selection)

	if err != nil {
		utils.ClearScreen()
		MainMenu(db)
	}

	if selection == 1 {
		utils.ClearScreen()
		MainMenu(db)
	} else if selection == 2 {
		os.Exit(1)
	}

	utils.ClearScreen()
	MainMenu(db)

}

func DeleteTransaction(db *gorm.DB) {
	var transactionName string

	fmt.Print("Type the transaction ID to be removed: ")

	_, err := fmt.Scanln(&transactionName)

	if err != nil {
		DeleteTransaction(db)
	}

	db.Delete(&structs.Transaction{}, transactionName)

	fmt.Println("Transaction successfully removed!")

	utils.ClearScreen()
	MainMenu(db)
}

func UpdateTransaction(db *gorm.DB) {
	var transactionName string
	var dbObject structs.Transaction
	fmt.Print("Type the transaction ID to be updated: ")

	_, err := fmt.Scanln(&transactionName)

	if err != nil {
		UpdateTransaction(db)
	}

	db.First(&dbObject)

	if err != nil {
		_ = fmt.Errorf("error during update")
	}

	fmt.Println("Wich field you want to update? \n 1 - Transaction name\n 2 - Transaction value")

	_, err = fmt.Scanln(&selection)

	if err != nil {
		UpdateTransaction(db)
	}

	if selection == 1 {
		var newName string
		fmt.Println("Type the new name: ")

		_, err = fmt.Scanln(&newName)

		if err != nil {
			UpdateTransaction(db)
		}

		dbObject.TransactionName = newName
		db.Save(&dbObject)

		fmt.Println("Transaction successfully updated!")

		fmt.Println("Go back to the menu? \n 1 - Yes\n 2 - No")

		_, err = fmt.Scanln(&selection)

		if err != nil {
			utils.ClearScreen()
			MainMenu(db)
		}

		if selection == 1 {
			utils.ClearScreen()
			MainMenu(db)
		} else if selection == 2 {
			os.Exit(1)
		}
	} else if selection == 2 {

		var newValue float64
		fmt.Println("Type the new value: ")

		_, err = fmt.Scanln(&newValue)

		if err != nil {
			UpdateTransaction(db)
		}

		dbObject.Value = newValue
		db.Save(&dbObject)

		fmt.Println("Transaction successfully updated!")

		fmt.Println("Go back to the menu? \n 1 - Yes\n 2 - No")

		_, err = fmt.Scanln(&selection)

		if err != nil {
			utils.ClearScreen()
			MainMenu(db)
		}

		if selection == 1 {
			utils.ClearScreen()
			MainMenu(db)
		} else if selection == 2 {
			os.Exit(1)
		}
	} else {
		UpdateTransaction(db)
	}

}
