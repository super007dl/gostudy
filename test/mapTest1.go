package main

import (
	"fmt"
)

// PersonInfo 个人信息类型
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {

	//初始化map
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	//map塞数据
	personDB["123"] = PersonInfo{"123", "join", "china"}
	personDB["1"] = PersonInfo{"1", "jack", "Japan"}

	//获取信息
	person, ok := personDB["123"]
	// ok 是bool，true表示找到值
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}


}
