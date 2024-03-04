package main

import "math"

func main() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)

	// var c float64 = 12.3
	// println(a + c) // 编译不通过

	// var d int32 = 50
	// println(a + d) // 编译不通过

	Bool()
	Byte()
	String()
	Extremum()
}

// Extremum 极值
func Extremum() {
	println("float64 最大值: ", math.MaxFloat64)
	// float64 没有最小值
	println("float64 最小的正数: ", math.SmallestNonzeroFloat64)

	println("float32 最大值: ", math.MaxFloat32)
	// float32 没有最小值
	println("float32 最小的正数: ", math.SmallestNonzeroFloat32)

	// int 族和 uint 族都有最大值最小值
}
