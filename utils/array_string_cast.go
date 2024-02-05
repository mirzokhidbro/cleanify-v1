package utils

import (
	"fmt"
	"strings"
)

func SetArray(input []string) string {
	replaced := strings.ReplaceAll(strings.Trim(fmt.Sprint(input), "[]"), " ", ", ")
	return "{" + replaced + "}"
}

func GetArray(input string) []string {
	data := strings.Trim(input, "{}")
	array := strings.Split(data, ",")
	return array
}
