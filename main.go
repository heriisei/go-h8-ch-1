package main

import "fmt"

func main() {
	i := 21
	j := true
	russia := 'Я'
	var k float64 = 123.456

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n\n", j)
	fmt.Printf("%b \n", i)
	fmt.Printf("%v \n", "Я")
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%c \n", 0x66)
	fmt.Printf("%c \n", 0x46)
	fmt.Printf("%U \n\n", russia)
	fmt.Printf("%f \n", k)
	fmt.Printf("%e \n", k)
}
