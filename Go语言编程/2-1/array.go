// array.go
package main

import (
	"fmt"
)

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(), array values:", array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)

	fmt.Println("In main(), array values:", array)

	for i := 0; i < len(array); i++ {
		fmt.Println("Element", i, "of array is", array[i])
	}

	for i, v := range array {
		fmt.Println("array element[", i, "]=", v)
	}
}
