package main

import "unicode/utf8"

func String() {
	// He said: "Hello, Go!"
	println("He said: \"Hello, Go!\"")
	println(`我可以换行
	这是新的行
	但是这里不能有反引号
	`)

	// 字符串只能和字符串拼接
	println("Hello, " + "Go!")

	// strings 包里面放着各种字符串相关操作的方法，需要的时候再查阅
	// strings.

	println(len("你好"))                      // 输出 6
	println(utf8.RuneCountInString("你好"))   // 输出 2
	println(utf8.RuneCountInString("你好ab")) // 输出 4
}
