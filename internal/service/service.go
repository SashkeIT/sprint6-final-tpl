package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Morse(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '/' {
			return false
		}
	}

	containsMorse := false
	for _, r := range s {
		if r == '.' || r == '-' {
			containsMorse = true
			break
		}
	}

	return containsMorse
}

func ToggleMorse(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	if Morse(s) {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
