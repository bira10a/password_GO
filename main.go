package main

import (
	"fmt"
	"password/bira10a/password_GO.git/account"
)

func main() {
	login := promptData("Введите login: ")
	password := promptData("Введите password: ")
	url := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Кривой у Вас URL")
		return
	}

	myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
