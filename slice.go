package awesomeProject1

import "fmt"

func AddPlus(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
	return slice
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(slice)

	fmt.Println(AddPlus(slice))
}
