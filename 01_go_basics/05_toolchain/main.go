package main

import (
	"fmt"

	"github.com/fatih/color"
)

func Add(a int, b int) int{
	return a+b
}

func Subtract(a int, b int) int {
	return a- b
}

func main(){
	color.Cyan("This is a cyan colored text")
	color.Green("This is a green colored text")
	color.Red("This is a red colored text")
	color.Blue("This is a blue colored text")
	color.HiBlue("This is a high blue colored text")
	fmt.Println("Addition : ",Add(1,2))
	fmt.Println("Subtraction : ",Subtract(1,2))
}