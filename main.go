package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

var letterRunes = []rune("1234567890qwertyuiopQWERTYUIOP")

func main() {
	fmt.Println(generatePassword(10))

	login := promptData("Введите login: ")
	password := promptData("Введите password: ")
	url := promptData("Введите url: ")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println((*acc).login, acc.password, acc.url)
}

func generatePassword(n int) string {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	return string(res)
}
