// slice.go
package main

import (
	"fmt"
)

func main() {
	var myArrayy [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArrayy[:5]

	mySlice1 := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	mySlice3 := []int{1, 2, 3, 4, 5}

	fmt.Println("Element of myArray:")
	for _, v := range myArrayy {
		fmt.Print(v, " ")
	}

	fmt.Println("\n")
	fmt.Println("Element of mySlice:")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println("\n")
	fmt.Println("mySlice1=", mySlice1)
	fmt.Println("mySlice2=", mySlice2)
	fmt.Println("mySlice3=", mySlice3)

	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))

	mySlice2 = append(mySlice2, 1, 2, 3)
	fmt.Println("mySlice2=", mySlice2)
	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))

	mySlice2 = append(mySlice2, mySlice3...)
	fmt.Println("mySlice2=", mySlice2)
	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))

	newSlice := mySlice3[:3]
	fmt.Println("newSlice=", newSlice)
}
