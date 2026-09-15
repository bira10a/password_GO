package main

import (
	"fmt"
	"password/bira10a/password_GO.git/account"
	"password/bira10a/password_GO.git/files"
)

func main() {
	fmt.Println("___Менеджер паролей___")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func findAccount() {

}

func deleteAccount() {

}

func createAccount() {
	// files.WriteFiles("Hello! I am Joi", "file.txt")
	files.ReadFiles("file.txt")

	login := promptData("Введите login: ")
	password := promptData("Введите password: ")
	url := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Кривой у Вас URL")
		return
	}

	file, err := myAccount.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать json")
		return
	}
	files.WriteFiles(file, "data.json")
	// myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
