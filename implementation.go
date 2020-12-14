package lab2

// PostfixToPrefix ...
// TODO: document this function
// PostfixToPrefix converts
func PostfixToPrefix(elem string) string {
	if include(elem) == 0 {
		return "Incorrect statement..."
	}
	if len(createArraySpace(elem, " ")) < 3 {
		return "dont enought args..."
	}
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
	input := " .0123456789+-*/^()"
	for j := 0; j < len(input); j++ {
		if char == string(input[j]) {
			return 1
		}
	}
	return 0
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
