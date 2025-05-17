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

func IntSliceToInterface(slice []int) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func Int8SliceToInterface(slice []int8) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func InterfaceSliceToInt(slice []interface{}) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		switch val := v.(type) {
		case int:
			result[i] = val
		case float64:
			result[i] = int(val)
		case string:
			// Try to parse string to int if needed
			var num int
			_, err := fmt.Sscanf(val, "%d", &num)
			if err != nil {
				return nil
			}
			result[i] = num
		default:
			return nil
		}
	}
	return result
}
