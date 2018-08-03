package main

import (
	"fmt"
)

var name string = "Isuru"

func main() {
	//fmt.Println("Hi")
	fmt.Println("fisrst :", name)
	//fmt.Println("1+1 =", 1.0+1.0)
	//fmt.Println(math.Sqrt(4))
	//foo()
	loop()
}
func foo() {
	fmt.Println("second :", name)
}

func loop() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func loop1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
