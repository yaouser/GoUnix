//统计核心数和任务数
package main

import(
	"fmt"
	"time"
	"runtime"
)

func main(){
	fmt.Println("CPU number:",runtime.NumCPU())
	fmt.Println("Goroutines start:",runtime.NumGoroutine())
	for i:=0;i<5;i++{
		go func(n int){
			fmt.Println(n,runtime.NumGoroutine())
		}(i)
	}
	time.Sleep(5*time.Second)
	fmt.Println("Goroutines over:",runtime.NumGoroutine())
}