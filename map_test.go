package golanggoroutine

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	index := strconv.Itoa(value)
	data.Store(index, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
	fmt.Println(data.Load("1")) // mendapatkan datanya secara tunggal

}
