// @Title  main
// @Description  自增整数生成器
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-09 23:30
package main

import "fmt"

// @title    IntegerGenerator
// @description   用于生成自增的整数
// @auth      MGAronya（张健）             2022-10-09 23:30
// @param     void
// @return    chan int			返回一个生成自增整数的管道
func IntegerGenerator() chan int {
	var ch chan int = make(chan int)

	// TODO 开启goroutine
	go func() {
		for i := 0; ; i++ {
			// TODO 直到通道索要数据才把i添加进入管道
			ch <- i
		}
	}()

	return ch
}

// @title    main
// @description   用于测试自增整数生成器
// @auth      MGAronya（张健）             2022-10-08 19:26
// @param     void
// @return    void
func main() {
	generator := IntegerGenerator()

	// TODO 生成100个自增的整数
	for i := 0; i < 100; i++ {
		fmt.Println(<-generator)
	}
}
