package main

import "fmt"

func format(items ...any) string {
	result := ""
	for _, item := range items {
		switch v := item.(type) {
		case int:
			result = fmt.Sprintf("%v %d", result, v)
		case float64:
			result = fmt.Sprintf("%v %f", result, v)
		case string:
			result = fmt.Sprintf("%v %q", result, v)
		case bool:
			if v {
				result = fmt.Sprintf("%v %s", result, "да")
			} else {
				result = fmt.Sprintf("%v %s", result, "нет")
			}
		default:
			result = fmt.Sprintf("%v %v", result, v)
		}
	}
	if result == "" {
		return result
	}
	return result[1:]
}

func main() {
	fmt.Println(format(42, 3.14159, "hi", true))
}
