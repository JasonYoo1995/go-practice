package main

import "fmt"

// 복합 타입

func main() {
	// array()
	// slice()
	// capacity()
	// makePractice()
	// slicing()
	// slicingWithAppend()
	slicingWithAppend2()
}

func array() {
	var x = [3]int{10, 20}
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2]) // 0으로 초기화

	var y = [...]int{10, 20} // 크기 자동 인식
	fmt.Println(y[0])
	fmt.Println(y[1])

	z := [...]int{10, 20, 0}
	fmt.Println(x == z) // 배열 간 값 비교 1
	z[2] = 30
	fmt.Println(x != z) // 배열 간 값 비교 2

	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Println(len(z))
}

func slice() {
	var x = []int{10, 20, 30}
	fmt.Println(x[1])

	var y []int
	fmt.Println(y == nil)
	fmt.Println(len(y))

	y = append(y, 1)
	y = append(y, 2, 3, 4)
	fmt.Println(y[0])
	fmt.Println(y[1])
	fmt.Println(y[2])
	fmt.Println(y[3])

	y = append(y, x...)
	fmt.Println(y[4])
	fmt.Println(y[5])
	fmt.Println(y[6])
}

func capacity() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func makePractice() {
	x := make([]int, 5) // 길이와 수용력이 모두 5 (5개의 값이 모두 0으로 초기화)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 123) // 길이 = 6, 수용력은 2배인 10
	fmt.Println(x, len(x), cap(x))

	y := make([]int, 0, 10)   // 길이는 0, 수용력은 10
	y = append(y, 1, 2, 3, 4) // 길이 = 4, 수용력은 10 유지
	fmt.Println(y, len(y), cap(y))
}

func slicing() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// 슬라이싱은 메모리를 공유한다
	x[1] = 20
	fmt.Println(x, y, z)
	y[0] = 10
	fmt.Println(x, y, z)
	z[1] = 30
	fmt.Println(x, y, z)
}

func slicingWithAppend() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(x, y)
	fmt.Println(len(x), len(y))
	fmt.Println(cap(x), cap(y))

	y = append(y, 30)
	fmt.Println(x, y)
	fmt.Println(len(x), len(y))
	fmt.Println(cap(x), cap(y))
}

func slicingWithAppend2() {
	x := make([]int, 0, 5)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 1, 2, 3, 4)
	fmt.Println(x, len(x), cap(x))

	y := x[:2] // y[4]에 해당하는 메모리 위치는 다른 슬라이스에 공유하지 않음
	z := x[2:] // z[2]에 해당하는 메모리 위치는 다른 슬라이스에 공유하지 않음
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))
	fmt.Println("비교 : ", x[2], z[0])
	fmt.Println("비교 : ", x[3], z[1])

	y = append(y, 30, 40, 50)
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))

	x = append(x, 60)
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))

	z = append(z, 70)
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))
	fmt.Println("비교 : ", x[4], y[4], z[2])
}
