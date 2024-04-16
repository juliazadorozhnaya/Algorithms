package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, char := range s {
		// если текущая скобка char - открывающая
		if closing, ok := matchingBrackets[char]; ok {
			stack = append(stack, closing)
		} else {
			// если char - закрывающая скобка, то должна соответствовать последнему элементу стека
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			// если match есть, удаляем последний элемент стека
			stack = stack[:len(stack)-1]
		}
	}
	// Стек должен быть пустым, если все скобки сбалансированы
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
