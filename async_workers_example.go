package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	goroutinesNum = 3
)


func startWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Println(strconv.Itoa(in), j)
	}
}

func main() {
	log.Print("Starting...")
	wg := &sync.WaitGroup{} // wait_2.go инициализируем группу
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) // добавляем воркер
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
	log.Println("Done")
}
