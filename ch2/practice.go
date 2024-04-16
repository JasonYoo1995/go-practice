package main

import "fmt"

func main() {
	// stringPractice()
	// typeConverionPractice()
	// declarePractice()
	constPractice()
}

func stringPractice() { // string 연습
	var my_str string = "hello"
	var my_str2 string = "hello2" // 타입 명시 O
	var my_str2_same = "hello2"   // 타입 명시 X
	var my_str_empty string       // default : 빈 값
	fmt.Println(my_str)
	fmt.Println(my_str < my_str2)
	fmt.Println(my_str > my_str2)
	fmt.Println(my_str == my_str2)       // 값 비교
	fmt.Println(my_str != my_str2)       // 값 비교
	fmt.Println(my_str2 == my_str2_same) // 값 비교
	fmt.Println(my_str_empty)            // 빈 값
	fmt.Println(my_str_empty == "")      // 빈 값인지 확인

	my_str += "3" // concatenate
	fmt.Println(my_str)

	my_str = "new" // 재할당
	fmt.Println(my_str)
}

func typeConverionPractice() { // 타입 변환 연습
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y) // 소수점 버림
	fmt.Println(z)
	fmt.Println(d)

	var y2 float64 = 30.9
	var d2 int = x + int(y2) // 소수점 버림
	fmt.Println(d2)
}

func declarePractice() { // 타입 선언 연습
	var x, y = 10, "hello"
	fmt.Println(x)
	fmt.Println(y)

	var (
		z    = 20
		d, r = "hello2", 'r'
		zero int
	)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(r)
	fmt.Println(zero)

	another := 123
	ano, ther := 1, 23
	fmt.Println(another)
	fmt.Println(ano)
	fmt.Println(ther)

	another, other := 3000, 4000 // another 재할당
	fmt.Println(another)
	fmt.Println(other)
}

func constPractice() { // 상수 연습
	var xVar = 3
	fmt.Println(xVar)
	xVar += 1
	fmt.Println(xVar)

	const xConst = 3
	fmt.Println(xConst)
	// xConst += 1 // 컴파일 에러 발생
	fmt.Print(xConst)
}
