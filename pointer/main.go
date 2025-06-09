package main

import "fmt"

func main() {
	pointer := 54

	a := &pointer

	x := 5
	y := 10

	updatePointer(a)

	swap(&x, &y)
	num := 7

	increment(&num)

	arr := []int{1, 2, 3, 4, 5}

	doubleValues(arr)

	fmt.Println("Swap", x, y)
	fmt.Println("Increment", num)
	fmt.Println("Doubled", arr)
}

func updatePointer(l *int) {
	*l = *l * 2
}

func swap(x *int, y *int) {
	temp := *x

	*x = *y
	*y = temp
}

func increment(inc *int) int {
	*inc = *inc + 1
	return *inc
}

func doubleValues(arr []int) {

	for i := range arr {
		arr[i] *= 2
	}

}
