package main

import "fmt"

func moveZeroes(nums []int) {
	p := 0 //Инициализация переменной p, которая будет использоваться как индекс для вставки не нулевых элементов в начало массива.
	// По сути, p отслеживает местоположение, куда следующий не нулевой элемент должен быть перемещен.
	for i, _ := range nums {
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			/*Текущий не нулевой элемент (nums[i]) меняется местами с элементом в позиции p.
			В случае если между p и i есть нулевые элементы, этот шаг перемещает не нулевой элемент на первую доступную
			позицию для не нулевых элементов и одновременно перемещает ноль ближе к концу массива.*/
			p++ //Увеличение p на 1, таким образом, сдвигая указатель на следующую позицию для вставки не нулевого элемента.
		}
	}
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
