package utils

import (
	"fmt"
	"strings"
)

func removeDuplicates(input []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func SetArray(input []string) string {
	input = removeDuplicates(input)
	replaced := strings.ReplaceAll(strings.Trim(fmt.Sprint(input), "[]"), " ", ", ")
	return "{" + replaced + "}"
}

func GetArray(input string) []string {
	data := strings.Trim(input, "{}")
	array := strings.Split(data, ",")
	return array
}
