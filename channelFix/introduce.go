package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// 现在有一个问题，如何高效的判断 1-20万 哪些数是素数

// lowEfficiencyPrime 暴力法,感觉根号已经算是高效了
func lowEfficiencyPrime() int {
	sum := 2 // 1,2 已经是素数了
	for i := 3; i < 200000; i++ {
		flag := true
		for j := 2; j < int(math.Sqrt(float64(i)))+1; j++ { // i-1 耗时 3.7s(17985)， i/2+1 耗时1.9s(17985) sqrt+1耗时 17.09ms
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			sum += 1
		}
	}
	return sum
}

// highEfficiencyPrime 使用并行（多个cpu）+ 优化（开根号）
// 45-50ms，但是结果不对（17984），难道是最后一个没拿到？而且明显不如上面那个嘛（人家17ms）
func highEfficiencyPrime() int {
	sum := 2
	val := make(chan bool, 10)
	go func() {
		for {
			if <-val {
				sum += 1
			}
		}
	}()
	for i := 3; i < 200000; i++ {
		go func(i int) {
			for j := 2; j < int(math.Sqrt(float64(i)))+1; j++ {
				if i%j == 0 {
					return
				}
			}
			val <- true
		}(i)
	}
	defer close(val)
	return sum
}
func main() {
	//fmt.Println(runtime.NumCPU())//16
	runtime.GOMAXPROCS(8)
	start := time.Now()
	num := highEfficiencyPrime()
	cost := time.Since(start)
	fmt.Println("cost:", cost, "num:", num)

}
