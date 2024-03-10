package golanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T){

	channel := make(chan string)
	defer close(channel) // akan dieksekusi paling akhir

	go func()  {
		time.Sleep(2 * time.Second)
		channel <- "Fauzan Nur Hidayat" // step 1
		fmt.Println("Berhasil mengirim channel ke variabel") // step 3
	}()

	data := <- channel // step 2
	fmt.Println(data) // step 4
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Fauzan Nur Hidayat New"
}

func TestCreateChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// in out
// hanya bisa menerima data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Fauzan Nur Hidayat" //mengisi data ke sebuah channel
}
// hanya bisa mengeluarkan data 
func OnlyOut(channel <-chan string) {
	// time.Sleep(2 * time.Second)
	data := <- channel // yang menerima data dari channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	
	go OnlyIn(channel)
	go OnlyOut(channel)
	
	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // memberikan kapasitas defer sebanyak 3 data
	defer close(channel)

	go func ()  {
		channel <- "Fauzan"
		channel <- "Nur"
		channel <- "Hidayat"
	}()
	
	// channel <- "New"
	go func ()  {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		// fmt.Println(<- channel)
	}()
	
	// fmt.Println(<- channel)
	time.Sleep(5 * time.Second)
	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 10; i++	 {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	
	// akan mengecek apakah ada channel, jika ada maka akan diterima dan di print
	for data := range channel {
		fmt.Println("Menerima data ", data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	go GiveMeResponse(channel1)

	counter := 0
	// di dalam for, sebuah variabel data akan menangkap data yang diberikan dari channel. entah 1 data, ataupun banyak data. sebanyak inisiasi pada method GiveMeResponse()
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1 : ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2 : ", data)
			counter++
		}
		if counter == 3 {
			break
		}
	}
}

func TestSelectChannelWithDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	go GiveMeResponse(channel1)

	counter := 0
	// di dalam for, sebuah variabel data akan menangkap data yang diberikan dari channel. entah 1 data, ataupun banyak data. sebanyak inisiasi pada method GiveMeResponse()
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1 : ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2 : ", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 3 {
			break
		}
	}
}

