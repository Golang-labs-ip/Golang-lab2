package main

import (
	"fmt"
)

func main() {
	userData := "4 2 - 3 * 5 +"
	if include(userData) == 1 {
		fmt.Println(postfixToPrefix(userData))
	} else {
		fmt.Println("Incorrect statement")
	}
}
func postfixToPrefix(elem string) string {
	return prefix(infix(elem))
}

func include(state string) int {
	if len(state) == 0 {
		return 0
	}
	for i := 0; i < len(state); i++ {
		if haveItem(string(state[i])) == 0 {
			return 0
		}
	}
	return 1
}
func haveItem(char string) int {
	input := " 0123456789+-*/^()"
	for j := 0; j < len(input); j++ {
		if char == string(input[j]) {
			return 1
		}
	}
	return 0
}

func prepare(elem []string) string {

	cur := ""
	for i := 0; i < len(elem); i++ {
		if elem[i] == "-" {
			if i == 0 {
				cur += "0"
			}
			if elem[i-1] == "(" {
				cur += "0"
			}
		}
		cur += elem[i]
	}
	elem = createArraySpace(cur, "")
	for i := 0; i < len(elem); {
		if i+1 < len(elem) && elem[i] == string("-") {
			if elem[i+1] == string("-") {
				elem[i] = "+"
				elem[i+1] = ""
				elem = createArraySpace(createString(elem), "")
				i--
			} else if elem[i+1] == string("+") {
				elem[i] = "-"
				elem[i+1] = ""
				elem = createArraySpace(createString(elem), "")
				i--
			} else {
				i++
			}

		} else if i+1 < len(elem) && elem[i] == string("+") {

			if elem[i+1] == string("-") {
				elem[i] = "-"
				elem[i+1] = ""
				elem = createArraySpace(createString(elem), "")
				i--
			} else if elem[i+1] == string("+") {
				elem[i] = "+"
				elem[i+1] = ""
				elem = createArraySpace(createString(elem), "")
				i--

			} else {
				i++
			}

		} else if elem[i] == ")" {
			if getPriority(elem[i-1]) != 0 {
				elem[i-1] = ""
				elem = createArraySpace(createString(elem), "")
				i++
			} else {
				i++
			}
		} else {
			i++
		}
	}
	return createString(elem)
}

func createArraySpace(elem string, space string) (result []string) {

	input := string([]byte(elem))
	arrayItem := ""
	for i := 0; i < len(input); i++ {
		if space == string("") {
			arrayItem += string(input[i])
			result = append(result, arrayItem)
			arrayItem = ""
		} else {
			if string(input[i]) != space {
				arrayItem += string(input[i])
			} else {
				result = append(result, arrayItem)
				arrayItem = ""
			}
			if i == len(input)-1 {
				result = append(result, arrayItem)
			}
		}
	}
	return
}

func createString(elem []string) (res string) {
	for _, v := range elem {
		res += v
	}
	return
}

func reverse(str string) (result string) {
	for _, v := range str {
		if string(v) == string("(") {
			result = ")" + result
		} else if string(v) == string(")") {
			result = "(" + result
		} else {
			result = string(v) + result
		}

	}
	return
}

func getPriority(elem string) int {

	if elem == string("^") {
		return 4
	} else if elem == string("*") || elem == string("/") {
		return 3
	} else if elem == string("+") || elem == string("-") {
		return 2
	} else if elem == string("(") {
		return 1
	} else if elem == string(")") {
		return -1
	} else {
		return 0
	}
}

func postfix(elem string) string {
	current := ""
	stack := ""
	for i := 0; i < len(elem); i++ {
		if getPriority(string(elem[i])) == 0 {
			current += string(elem[i])
		}
		if getPriority(string(elem[i])) == 1 {
			stack += string(elem[i])
		}
		if getPriority(string(elem[i])) == -1 {
			for j := len(stack) - 1; j >= 0; j-- {
				if getPriority(string(stack[j])) != 1 {
					current += " "
					current += string(stack[j])
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
					break
				}
			}
		}
		if getPriority(string(elem[i])) > 1 {
			current += " "
			for j := len(stack) - 1; j >= 0; j-- {
				if getPriority(string(stack[j])) >= getPriority(string(elem[i])) {
					current += string(stack[j])
					current += " "
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack += string(elem[i])
		}
	}
	for j := len(stack) - 1; j >= 0; j-- {
		current += " "
		current += string(stack[j])
	}
	return current
}

func prefix(elem string) string {
	return reverse(postfix(reverse(elem)))
}

func infix(elem string) string {
	array := createArraySpace(elem, " ")
	stack := make([]string, 0)
	str := ""
	for i := 0; i < len(array); i++ {
		if getPriority(array[i]) == 0 {
			stack = append(stack, array[i])
		} else {
			length := len(stack)
			str += "("
			str += stack[len(stack)-2]
			str += array[i]
			str += stack[len(stack)-1]
			str += ")"
			stack = stack[0 : length-2]
			stack = append(stack, str)
			str = ""
		}
	}

	return createString(stack)
}
