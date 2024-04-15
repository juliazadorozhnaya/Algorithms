package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	readIndex, writeIndex := 0, 0
	/*readIndex для чтения текущего символа из массива chars
	writeIndex для записи символов в массив chars*/
	for readIndex < len(chars) {
		/*Цикл продолжается до тех пор, пока readIndex не достигнет конца массива chars.*/
		currentChar := chars[readIndex]
		currentCount := 0
		/*  Выбирается текущий символ и подсчитывается, сколько раз он подряд встречается в массиве.
		currentChar хранит текущий символ, а currentCount - количество его последовательных повторений.*/
		for readIndex < len(chars) && currentChar == chars[readIndex] {
			readIndex++
			currentCount++
		}

		/*
			После подсчета повторений текущий символ записывается в массив на позицию writeIndex, после чего writeIndex
			инкрементируется. Если символ встречался более одного раза (currentCount > 1), его количество преобразуется
			в строку с помощью strconv.Itoa(currentCount), и каждая цифра этого числа записывается в массив как отдельный символ.
		*/

		chars[writeIndex] = currentChar
		writeIndex++

		if currentCount > 1 {
			count := strconv.Itoa(currentCount)
			for _, i := range count {
				chars[writeIndex] = byte(i)
				writeIndex++
			}
		}
	}
	return writeIndex //Функция возвращает новую длину массива chars, которая теперь равна writeIndex.
}

func main() {
	// Входной массив символов
	chars := []byte("aabbbccc")
	fmt.Println("Before compression:", string(chars))

	// Вызов функции сжатия и получение новой длины массива
	length := compress(chars)
	fmt.Println("After compression:", string(chars[:length]))
}
