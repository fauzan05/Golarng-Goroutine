package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	// go fmt.Println("Hello guyss1")
	fmt.Println("Hello guyss2")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number : ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i <= 1000000; i++ {
		go DisplayNumber(i)
	}
	// time.Sleep(20 * time.Second)
}