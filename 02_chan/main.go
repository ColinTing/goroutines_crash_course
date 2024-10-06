package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < 100; i++ {
			fmt.Println("第一个IDX生成器:", i)
			sum += i
		}
		// 将变量 sum 的值发送到通道 c 中
		c <- sum
	}()
	// 接收操作：从通道中取出数据，并且将这个数据赋值给一个变量
	output := <-c
	fmt.Println("第一个IDX生成器的总和:", output)
}
