package main

import "fmt"

/**
*数组切片 弥补数组，只能初始化定义长度，不可更改问题
 */
func main() {
	//定义并声明数组
	var myArray = [10] int{1,2,3,4,5,6,7,8,9,10}

	//基于数组创建数组切片 myArray[first:last]
	var mySlice = myArray [:5]

	//打印数组的元素
	fmt.Println("myArray Elements:")
	for _, v := range myArray{
		fmt.Print(v, "  ")
	}

	//打印切片的元素
	fmt.Println("mySlice Elements:")
	for _,v := range mySlice{
		fmt.Print(v, "  ")
	}

	//默认所有元素
	var mySlice1 = myArray [:]
	//打印切片的元素
	fmt.Println("mySlice1 Elements:")
	for _,v := range mySlice1{
		fmt.Print(v, "  ")
	}

	//默认从第5个元素开始
	var mySlice2 = myArray [5:]
	//打印切片的元素
	fmt.Println("mySlice2 Elements:")
	for _,v := range mySlice2{
		fmt.Print(v, "  ")
	}


	makeTest()

}

//通过make方法创建切片
func makeTest()  {
	//创建一个初始元素个数为5的数组切片，元素初始值为0
	mySlice1 := make([]int, 5)

	//创建一个初始元素个数为5的数组切片，元素初始值为0，
	//并预留10个元 的存储空间
	mySlice2 := make([]int, 5, 10)


	//直接创建并初始化包含7个元素的数组切片:
	mySlice3 := []int{1, 2, 3, 4, 5, 6, 7}

	//数组遍历方式1.同其他语言
	for i := 0; i <len(mySlice1); i++ {
		fmt.Println("mySlice1[", i, "] =", mySlice1[i])
	}


	//cap()函数返回的是数组切片分配的空间大小，而len()函数返回的是 数组切片中当前所存储的元素个数
	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))

	mySlice2 = append(mySlice3)

	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))
	for i := 0; i <len(mySlice2); i++ {
		fmt.Println("mySlice2[", i, "] =", mySlice2[i])
	}

	//使用range关键字可以遍历代码显得更整洁。range表达式有两个返回值，第一个是索引引，
	//第二个是元素的值
	for i, v := range mySlice3{
		fmt.Print("mySlice3[", i, "] =", v)
	}

}