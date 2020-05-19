package main

import "fmt"

func main() {

	fmt.Println("hello world ……")

	/**
	定义变量
	*/

	// 第一种
	var num int = 10
	fmt.Printf("数值：%d\n", num)

	// 第二种
	var name = "王二狗"
	fmt.Printf("类型是：%T，数值是%s\n", name, name)

	// 第三种
	sex := "男"
	fmt.Printf("sex:%s", sex)

}
