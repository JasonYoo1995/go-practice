package main

import "fmt"

func main() {
	title_6_1()
	title_6_2()
	title_6_3()
}

func title_6_1() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX) // x의 메모리 주소를 출력

	fmt.Println(*pointerToX) // x의 값을 출력

	z := 5 + *pointerToX
	fmt.Println(z)
}

func title_6_2() {
	x := 10
	y := x
	y = 20
	fmt.Println(x, y)
}

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func title_6_3() {
	x := 10

	failedUpdate(&x)
	fmt.Println(x)

	update(&x)
	fmt.Println(x)
}
