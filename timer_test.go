package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C // mengembalikan data time lewat channel
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second) // langsung mengembalikan datanya berupa channel
	fmt.Println(time.Now())

	time := <- channel // membawa data time terkini setelah 5 detik
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	fmt.Println(time.Now())
	group.Add(1)
	time.AfterFunc(5 * time.Second, func ()  {
		fmt.Println(time.Now())
		fmt.Println("Execute after 5 second")
		group.Done()
	})

	group.Wait()
}