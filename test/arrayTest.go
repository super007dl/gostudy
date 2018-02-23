package main

import "fmt"

/**
所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作
函数体中无法 改传入的数组的内容，因为函数内操作的只是所 传入数组的一个副本。
 */
func main() {
	//定义并初始化数组
	array := [5]int{1,2,3,4,5}

	//方法内部修改数组
	modify(array)

	fmt.Println("In main(), array values:", array)

}

func modify(array [5]int)  {
	array[0] = 10
	fmt.Println("In modify(), array values:", array)

}
