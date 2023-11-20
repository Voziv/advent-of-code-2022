package main

import "fmt"

type Data struct {
	A int
}

func ExampleCopy(d Data) *Data {
	d.A = 2
	return &d
}

func ExampleCopyTwo(d Data) Data {
	d.A = 2
	return d
}

func ExamplePointer(d *Data) {
	d.A = 3
}

func main() {
	myOriginalData := Data{A: 69}

	fmt.Println(myOriginalData)

	newData := ExampleCopy(myOriginalData)
	fmt.Println(newData)
	fmt.Println(myOriginalData)

	ExamplePointer(&myOriginalData)
	fmt.Println(newData)
	fmt.Println(myOriginalData)
}
