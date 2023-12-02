package service

import (
	"errors"
	"math/rand"
	"net/url"
)

var (
	ErrUrlAlreadyExists = errors.New("url already exists in storage")
	ErrUrlDoesNotExist  = errors.New("url doesn't exist")
)

const aliasLength = 10

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

func GenerateAlias() string {
	b := make([]rune, aliasLength)

	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}

func IsValid(urlToValidate string) bool {
	_, err := url.ParseRequestURI(urlToValidate)
	if err != nil {
		return false
	}
	return true
}
