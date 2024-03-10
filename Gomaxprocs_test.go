package golanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func ()  {
			// fmt.Println(i)	
		}()
		group.Done()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCpu)

	// runtime.GOMAXPROCS(20) // menetapkan jumlah thread yang akan digunakan
	totalThread := runtime.GOMAXPROCS(0)
	fmt.Println("Total Thread : ",totalThread)

	group.Wait()
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)
}