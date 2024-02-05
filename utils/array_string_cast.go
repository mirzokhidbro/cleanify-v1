package utils

import (
	"fmt"
	"strings"
)

func SetArray(input []string) string {
	replaced := strings.ReplaceAll(strings.Trim(fmt.Sprint(input), "[]"), " ", ", ")
	return "{" + replaced + "}"
}
