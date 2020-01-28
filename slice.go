package main

import "fmt"

func AddPlus(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
	return slice
}
func Add5(slice []int) []int {
	return append(slice, 5)
}
func StartAdd5(slice []int) []int {
	return append([]int{5}, slice...)
}
func Last(slice []int) (int, []int) {
	last := slice[len(slice)-1]
	return last, slice[:len(slice)-1]
}
func First(slice []int) (int, []int) {
	first := slice[0]
	return first, slice[1:]
}

func Popik(slice []int, i int) (int, []int) {
	value := slice[i]
	copy(slice[i:], slice[i+1:])
	return value, slice
}
func ArrPlusArr(slice1, slice2 []int) []int {
	return append(slice1, slice2...)
}
func ShiftLeft(slice []int) []int {
	first := slice[0]
	copy(slice, slice[1:])
	slice[len(slice)-1] = first
	return slice
}
func ShiftLeftI(slice []int, shift int) []int {
	var result = make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		j := i - shift
		if j < 0 {
			j += len(slice)
		}
		result[j] = slice[i]
	}
	return result
}

func ShiftRightI(slice []int, shift int) []int {
	return ShiftLeftI(slice, shift)
}

func main() {
	slice := []int{0, 1, 2, 3, 4}
	slice1 := []int{5, 6, 7, 8, 9}

	fmt.Println(slice)

	fmt.Println(AddPlus(slice))
	fmt.Println(Add5(slice))
	fmt.Println(StartAdd5(slice))
	fmt.Println(Last(slice))
	fmt.Println(First(slice))
	fmt.Println(Popik(slice, 2))
	fmt.Println(ArrPlusArr(slice, slice1))
	fmt.Println(ShiftLeft(slice1))
	fmt.Println(ShiftLeftI(slice1, 4))
	fmt.Println(ShiftRightI(slice1, 2))

}
