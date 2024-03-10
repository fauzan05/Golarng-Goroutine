package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T){
	var once sync.Once = sync.Once{}
	var group sync.WaitGroup = sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func ()  {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}