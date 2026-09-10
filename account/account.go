package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var letterRunes = []rune("1234567890qwertyuiopQWERTYUIOP")

type account struct {
	login    string
	password string
	url      string
}

func (acc account) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
}

func NewAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAccount := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if password == "" {
		newAccount.generatePassword(10)
	}

	return newAccount, nil
}
