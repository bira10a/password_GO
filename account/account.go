package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("1234567890qwertyuiopQWERTYUIOP")

type account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreateAt  time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc account) OutputPassword() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.Password = string(res)
}

func (acc *account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)

	if err != nil {
		return nil, err
	}
	return file, nil
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
