package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

func readFromKeyboard() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	print("Please enter an expression with either only addition, subtraction, or multiplication. Please put a space in between each of the numbers and operations. At the end of your equation, press Ctrl + Z, then click enter. Equation:")
	y := stack.New()
	x := stack.New()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {

		str := scanner.Text()
		if str == "+" || str == "-" || str == "*" {
			y.Push(str)
		} else {
			i, err := strconv.Atoi(str)
			if err != nil {
				print("YOU HAVE RECIEVED AN F IN READING, COUNTING, COMMON SENSE, AND WHATNOT... PLEASE GO TO A DOCTOR!!!")
				return
			}
			x.Push(i)

		}
		if x.Len() == y.Len()+1 && x.Len() > 1 {
			a := x.Pop().(int)
			b := y.Pop().(string)
			c := x.Pop().(int)
			d := 0
			if b == "-" {
				d = c - a
			} else if b == "+" {
				d = a + c
			} else {
				d = a * c
			}
			x.Push(d)
		}
	}

	ans := x.Pop().(int)
	print("The answer to your equation is ", ans)
	// intgr := 0
	// fmt.Scanf("%d", &intgr)
	// op := ' '
	// fmt.Scanf("%c", &op)
	// fmt.Scanf("\n")
	// fmt.Println(intgr, op)
}
