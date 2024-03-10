package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T){
	pool := sync.Pool{
		New: func() any {
			return "Tidak kebagian"
		},
	}

	// data1 := "Satu"
	// data2 := "Dua"
	// data3 := "Tiga"

	// pool.Put(&data1)
	// pool.Put(&data2)
	// pool.Put(&data3)

	pool.Put("Satu")
	pool.Put("Dua")
	pool.Put("Tiga")


	for i := 0; i < 10; i++ {
		go func ()  {
			// mengambil data dari pool
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(2 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("selesai")
}