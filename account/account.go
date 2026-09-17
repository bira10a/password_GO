package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("1234567890qwertyuiopQWERTYUIOP")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreateAt  time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) OutputPassword() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAccount := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		newAccount.generatePassword(10)
	}

	return newAccount, nil
}
