package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // akan mengirim data berupa channel setiap 1 detik
	// fmt.Println(ticker)

	go func ()  {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		// menerima data yang diberikan oleh channel ticker
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}