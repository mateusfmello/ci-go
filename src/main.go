package main

import "fmt"

func Soma(num1 float64, num2 float64) (float64) {

	return num1 + num2;
}


func main() {

	resultado := Soma(5,5);
	
	println(fmt.Sprintf("%g", resultado));
}