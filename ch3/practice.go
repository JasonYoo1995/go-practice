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
	// slicingWithout3rdArgument()
	// fmt.Println()
	// slicingWith3rdArgument()

	/* 여기부터 3.3*/
	// stringPractice()
	// conversion()
	// mapPractice()
	// ok()
	// mapSet()
	structPractice()
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

/* 여기부터 3.3*/

func stringPractice() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b) // t = 116

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s, "/", s2, "/", s3, "/", s4)

	s4 = "there2"
	fmt.Println(s, "/", s2, "/", s3, "/", s4)

	var ss string = "Hello 😍" // "Hello XXXX" 처럼 😍는 4바이트를 차지
	var ss2 string = ss[4:7]
	var ss3 string = ss[:5]
	var ss4 string = ss[6:]
	fmt.Println(ss, "/", ss2, "/", ss3, "/", ss4)

	fmt.Println(len(s), len(ss))

	var ss5 string = ss[4:10]
	fmt.Println(ss5)
}

func conversion() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s, s2)

	var ss string = "Hello, 😝"
	var bs []byte = []byte(ss)
	var rs []rune = []rune(ss)
	fmt.Println(bs)
	fmt.Println(rs)
	fmt.Println(len(bs), len(rs))
}

func mapPractice() {
	var nilMap map[string]int // write 불가 (패닉 발생)
	fmt.Println(nilMap)
	fmt.Println(nilMap["any"])

	totalWins := map[string]int{} // write 가능
	fmt.Println(totalWins)
	fmt.Println(totalWins["any"])
	totalWins["any"] = 1
	fmt.Println(totalWins["any"])

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	fmt.Println(teams["any"])
	fmt.Println(teams["orcas"])
	fmt.Println(teams["Orcas"])
	fmt.Println(teams["Lions"][0])
	fmt.Println(teams["Lions"][2])

	ages := make(map[int][]string, 3)
	ages[10] = []string{"배움", "순수", "성장"}
	ages[20] = []string{"젊음", "패기", "열정"}
	ages[30] = []string{"책임", "도전", "성숙"}
	ages[40] = []string{"건강", "요령", "리더십"}
	fmt.Println(ages)
}

func ok() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v2, ok := m["world"]
	fmt.Println(v2, ok)

	v3, ok := m["goodbye"]
	fmt.Println(v3, ok)
}

func mapSet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

func structPractice() {
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	g = f
	fmt.Println(f == g)

	g.name = "Tom"
	fmt.Println(f == g)

	anonymous := struct {
		name string
		age  int
	}{
		name: "Bob",
		age:  50,
	}
	fmt.Println(f == anonymous)

	fmt.Println(f, g, anonymous)
}
