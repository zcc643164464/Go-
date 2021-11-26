package main

import "fmt"
/*
当两个或多个操作必须按正确的顺序执行，而程序并未保证这个顺序，就会
发生竞争条件。
*/
func main(){
	var data int = 0
	go func() {
		data++
	}()
	if data == 0 {
		// 此行并不会打印
		fmt.Printf("the value is %v.\n",data)
	}
}