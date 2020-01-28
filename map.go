package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func CountWords(text string) map[string]int {
	split := strings.Split(text, " ")
	counts := make(map[string]int)

	for _, count := range split {
		counts[count]++
	}

	return counts
}

func Numeral(slice []int) map[int]int {
	num := make(map[int]int)

	for _, i := range slice {
		num[i]++
	}

	return num
}
func Matches(slice1 []int, slice2 []int) (result []int) {
	num1 := Numeral(slice1)
	num2 := Numeral(slice2)

	for k := range num1 {
		if _, ok := num2[k]; ok {
			result = append(result, k)
		}
	}

	return
}
func main() {
	text := "К каждому элементу []int прибавить 1 Добавить в конец слайса число 5Добавить в начало слайса число 5Взять последнее число слайса, вернуть его пользователю, а из слайса этот элемент удалитьВзять первое число слайса, вернуть его пользователю, а из слайса этот элемент удалитьВзять i-е число слайса, вернуть его пользователю, а из слайса этот элемент удалить. Число i передает пользователь в функцию"
	fmt.Println(CountWords(text))

	var slice = make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = rand.Intn(10)
	}
	fmt.Println(Numeral(slice))
	var slice2 = make([]int, 100)
	for i := 0; i < 100; i++ {
		slice2[i] = rand.Intn(10)
	}
	fmt.Println(Matches(slice, slice2))

}
