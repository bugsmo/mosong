package main

func Byte() {
	var a byte = 'a'
	println(a) // 输出的是 a 的 ASCII 表达 97

	var str string = "this is string"
	var bs []byte = []byte(str)
	println(bs)
}
