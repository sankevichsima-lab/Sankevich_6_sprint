package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func CheckMorse(text string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", fmt.Errorf("empty string")
	}

	isNotMorse := false
	for _, char := range text {
		if char != ' ' && char != '.' && char != '-' {
			isNotMorse = true
			break
		}
	}

	if isNotMorse {
		return morse.ToMorse(text), nil
	}

	return morse.ToText(text), nil
}
