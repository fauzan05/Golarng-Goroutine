package golanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x  int64 = 0
	group := sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for i := 1; i <= 100; i++ {
				// alternatif lain selain menggunakan mutex
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter : ", x)
}