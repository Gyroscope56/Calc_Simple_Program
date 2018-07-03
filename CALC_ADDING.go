package main

import (
	"fmt"
)

func main2() {
	println("There will be 15 numbers you can insert. If you have 14 or less, you can fill in the other numbers with 0. Please also note that this calculator can only add numbers together.")
	print("Number 1:")
	int1 := 0
	fmt.Scanf("%d\n", &int1)
	print("Number 2:")
	int2 := 0
	fmt.Scanf("%d\n", &int2)
	print("Number 3:")
	int3 := 0
	fmt.Scanf("%d\n", &int3)
	print("Number 4:")
	int4 := 0
	fmt.Scanf("%d\n", &int4)
	print("Number 5:")
	int5 := 0
	fmt.Scanf("%d\n", &int5)
	print("Number 6:")
	int6 := 0
	fmt.Scanf("%d\n", &int6)
	print("Number 7:")
	int7 := 0
	fmt.Scanf("%d\n", &int7)
	print("Number 8:")
	int8 := 0
	fmt.Scanf("%d\n", &int8)
	print("Number 9:")
	int9 := 0
	fmt.Scanf("%d\n", &int9)
	print("Number 10:")
	int10 := 0
	fmt.Scanf("%d\n", &int10)
	print("Number 11:")
	int11 := 0
	fmt.Scanf("%d\n", &int11)
	print("Number 12:")
	int12 := 0
	fmt.Scanf("%d\n", &int12)
	print("Number 13:")
	int13 := 0
	fmt.Scanf("%d\n", &int13)
	print("Number 14:")
	int14 := 0
	fmt.Scanf("%d\n", &int14)
	print("Number 15:")
	int15 := 0
	fmt.Scanf("%d\n", &int15)
	answer := int1 + int1 + int3 + int4 + int5 + int6 + int7 + int8 + int9 + int10 + int11 + int12 + int13 + int14 + int15
	println("The added sum of all of your numbers is", answer)
}
