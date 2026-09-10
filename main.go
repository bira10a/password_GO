package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAccount := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if login == "" {
		newAccount.generatePassword(10)
	}

	return newAccount, nil
}

var letterRunes = []rune("1234567890qwertyuiopQWERTYUIOP")

func main() {
	login := promptData("Введите login: ")
	password := promptData("Введите password: ")
	url := promptData("Введите url: ")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Кривой у Вас URL")
		return
	}

	myAccount.generatePassword(10)
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
