package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	done := make(chan bool)

	fmt.Println(len(done), cap(done))

	defer close(message)

	// consumer
	go func() {
		// time.Ticker(NewTicker是个工厂方法，创建Ticker对象)：按照指定时间向通道C发送消息
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			// select 语句是用来监听多个 channel 的消息，非常类似于switch
			select {
			// 如果 done channel 有消息，则执行其流程代码
			// 如果 done 已经关闭，则会立刻读到 0 值(因为是bool类型所以是false)
			case <-done:
				fmt.Println("child process interrupt", <-done)
				return
			// default是默认分支，当没有其他分支可以执行时，会执行default分支的代码（可选）
			default:
				fmt.Println("send message: ", <-message)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		message <- i
	}

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit")
}
