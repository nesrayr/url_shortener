package utils

import (
	"errors"
	"math/rand"
	"net/url"
)

const aliasLength = 10

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

func GenerateAlias(url string) string {
	b := make([]rune, aliasLength)

	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}

func IsValid(urlToValidate string) error {
	_, err := url.ParseRequestURI(urlToValidate)
	if err != nil {
		return errors.New("invalid url")
	}
	return nil
}
