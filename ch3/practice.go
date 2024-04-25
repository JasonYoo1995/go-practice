package main

import "fmt"

// 복합 타입

func main() {
	array()
	slice()
	capacity()
	makePractice()
	slicing()
	slicingWithAppend()
	slicingWithout3rdArgument()
	fmt.Println()
	slicingWith3rdArgument()
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

func slicingWithout3rdArgument() { // 원본 슬라이스(x)의 사용되지 않은 수용력은 모든 하위 슬라이스에 공유된다
	x := make([]int, 0, 5)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 1, 2, 3, 4)
	fmt.Println(x, len(x), cap(x))

	y := x[:2]
	z := x[2:]
	// x[0]는 y[0]에 공유된다
	// x[1]는 y[1]에 공유된다
	// x[2]는 y[2]와 z[0]에 공유된다
	// x[3]는 y[3]와 z[1]에 공유된다
	// x[4]는 y[4]에 공유된다

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

	// 새로 할당된 z[2]는 x[4]와 y[4]에 공유된다
	z = append(z, 70)
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))
	fmt.Println("비교 : ", x[4], y[4], z[2])

	x[0] = 10
	y[1] = 20
	fmt.Println(x, y, z)
}

// 완전한 슬라이스 연산 (Full Slice Expression)
//   - 용량이 안 바뀌면 공유
//   - 용량이 바뀌면 새로운 슬라이스로 deep copy
func slicingWith3rdArgument() {
	x := make([]int, 0, 5)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 1, 2, 3, 4)
	fmt.Println(x, len(x), cap(x))

	/*
		* x[a:b:c]에서 각 인자의 의미
		 - a: 시작 인덱스
		 - b: 끝 인덱스 (해당 인덱스는 포함되지 않음)
		 - c: 용량(capacity) 지정

		* 예시
		 - x[a:b]는 x 슬라이스의 인덱스 a부터 b-1까지를 선택합니다.
		 - x[a:b:c]는 x 슬라이스의 인덱스 a부터 b-1까지를 선택하며, 용량을 c-a로 설정합니다. 이렇게 함으로써 슬라이스의 길이와 용량을 제어할 수 있습니다.
	*/

	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))

	y[0] = 100 // 공유됨 (용량이 안 바뀐 상황)
	fmt.Println(x, y, z)

	y = append(y, 30, 40, 50) // 공유 안 됨 (용량이 바뀐 상황)
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))

	y[1] = 200 // 공유 안 됨 (용량이 바뀐 상황)
	fmt.Println(x, y, z)

	x = append(x, 60)
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))

	z = append(z, 70)
	fmt.Println(x, y, z)
	fmt.Println("len : ", len(x), len(y), len(z))
	fmt.Println("cap : ", cap(x), cap(y), cap(z))

	y[0] = 10
	fmt.Println(x, y, z)
}
