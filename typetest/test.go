package main

import "fmt"

type name string

func (n name) len() int {
	return len(n)
}

func main() {
	var myname name = "taozs" //其实就是字符串类型
	l := []byte(myname)       //字符串转字节数组
	fmt.Println(l)
	fmt.Println(len(l))       //字节长度
	fmt.Println(myname.len()) //调用对象的方法
}
