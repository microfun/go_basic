package main

import "fmt"

// 공개 함수 (대문자로 시작)
func Add(a int, b int) int {
	return a + b
}

// 비공개 함수 (소문자로 시작)
func subtract(a int, b int) int {
	return a - b
}

// 공개 구조체
type Person struct {
	Name string // 공개 필드
	age  int    // 비공개 필드
}

func main() {
	// 변수명은 소문자로 시작
	var result int
	result = Add(3, 4)
	fmt.Println("Addition:", result)

	// 비공개 함수 호출
	result = subtract(10, 5)
	fmt.Println("Subtraction:", result)
}