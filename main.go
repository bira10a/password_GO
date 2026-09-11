package main

import (
	"fmt"
	"password/bira10a/password_GO.git/account"
	"password/bira10a/password_GO.git/files"
)

func main() {

	files.WriteFiles("Hello! I am Joi", "file.txt")
	// files.ReadFiles()

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
