package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Morse(s string) bool {
	idx := strings.IndexFunc(s, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})
	return idx == -1
}

func ToggleMorse(s string) (string, error) {
	if Morse(s) {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
