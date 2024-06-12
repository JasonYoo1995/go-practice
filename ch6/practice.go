package main

import "fmt"

func main() {
	// title_6_1()
	// title_6_2()
	// title_6_3()
	title_6_4()
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

// *************************************************
/** 2024.6.12 (수) */

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

func MakeFoo2() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

func title_6_4() {
	var f Foo // initialize empty struct
	MakeFoo(&f)
	fmt.Println("f =", f)

	f2, _ := MakeFoo2()
	fmt.Println("f2 =", f2)

	/**

	[함수로 넘길 때]
	- 맵 : Shallow Copy
	- 구조체 : Deep Copy

	[가변성과 불변성]
	- Mutability = 가변성 = ex) Reference Type(참조 타입)
	- Immutability = 불변성 = ex) Primitive Type(원시 타입)
	- 불변성을 고수하는 것이 코드의 유지보수성과 가독성 측면에서 유리하다 (단, 성능 최적화 목적의 가변성은 허용)

	*/
}
