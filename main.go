package main

import "fmt"

type StringableFloat interface {
	~float32 | ~float64
	String() string
}

type myFloat float64

func (m myFloat) String() string {
	return fmt.Sprintf("%e", m)
}

func StringyfyFloat[T StringableFloat](f T) string {
	return f.String()
}

func main() {
	var f myFloat = 1.234
	s := StringyfyFloat(f)

	fmt.Println(s)
}
