package main

import (
	"fmt"
	"sort"
)

func main() {
	testSlice2()
}

func testArrays1() {
	var arr [3]string

	arr[0] = "GO"
	arr[1] = "NODE"
	arr[2] = "PYTHON"

	fmt.Println(arr)
}

// arr elon qilishning qisqa usuli
func testArrays2() {
	arr := [3]int{9, 4, 7}
	fmt.Println(arr)
}

func testArrays3() {
	arrs := [3][3]string{
		{"GO", "NODE", "PYTHON"},
		{"VUE", "REACT", "ANGULAR"},
		{"CSS", "SASS", "LESS"}}
	fmt.Println(arrs)
}

func testArrays4() {
	arr1 := [3]int{1, 2, 3}

	arr2 := &arr1
	fmt.Println(arr1)
	fmt.Println(*arr2)

	arr1[2] = 4
	fmt.Println(arr1)
	fmt.Println(*arr2)
}

// slice

func testSlice1() {
	// yengi slice elon qilish
	myslice := []int{2, 4, 8}

	// slice oxiriga yengi element qo'shish
	myslice = append(myslice, 10)

	// slice uzunligini olish
	fmt.Println(len(myslice))
}

func testSlice2() {
	myarr := [6]string{"GO", "NODE", "PHP", "C++", "C#", "PYTHON"}
	myslice := myarr[1:3]
	fmt.Println(myslice)
}

func testSlice3() {
	myslice := []int{45, 23, 75, 54, 14}
	sort.Ints(myslice)
	fmt.Println(myslice)
}

