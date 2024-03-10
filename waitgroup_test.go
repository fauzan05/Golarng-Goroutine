package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	// "time"
)

func RunAsyncronous(group *sync.WaitGroup, iteration int) {
	defer group.Done()

	// group.Add(1)

	fmt.Println("Hello ke : ", iteration)
	// time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T){
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go RunAsyncronous(group, i)
	}

	group.Wait()
	// time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
}