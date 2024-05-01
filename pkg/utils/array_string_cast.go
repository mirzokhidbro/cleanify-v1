package utils

import (
	"fmt"
	"strings"
)

func removeDuplicates(input []interface{}) []interface{} {
	keys := make(map[interface{}]bool)
	list := []interface{}{}

	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func SetArray(input []interface{}) interface{} {
	input = removeDuplicates(input)
	replaced := strings.ReplaceAll(strings.Trim(fmt.Sprint(input), "[]"), " ", ", ")
	return "{" + replaced + "}"
}

func GetArray(input interface{}) []interface{} {
	str, ok := input.(string)
	if !ok {
		return nil
	}
	data := strings.Trim(str, "{}")
	array := strings.Split(data, ",")
	result := make([]interface{}, len(array))
	for i, v := range array {
		result[i] = v
	}
	return result
}

func StringSliceToInterface(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func InterfaceSliceToString(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		str, ok := v.(string)
		if !ok {
			return nil
		}
		result[i] = str
	}
	return result
}

func IntSliceToInterface(slice []int8) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func InterfaceSliceToInt(slice []interface{}) []int8 {
	result := make([]int8, len(slice))
	for i, v := range slice {
		num, ok := v.(int8)
		if !ok {
			return nil
		}
		result[i] = num
	}
	return result
}
