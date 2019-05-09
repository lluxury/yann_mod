package yann_mod

import (
	"errors"
	"fmt"
)

func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	case "cn":
		return fmt.Sprintf("你好，%s！", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
