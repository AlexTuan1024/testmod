package testmod

import (
	"fmt"
	"github.com/pkg/errors"
)

// Hi returns a friendly greeting
func Hi(name string, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi,%s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi,%s!", name), nil
	case "es":
		return fmt.Sprintf("Hola,%s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour,%s!", name), nil
	case "cn":
		return fmt.Sprintf("你好，%s!", name), nil
	case "jp":
		return fmt.Sprintf("こんにちは、%s!", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
