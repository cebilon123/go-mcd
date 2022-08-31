package main

import (
	"log"
	"sync"
	"time"
)

const (
	timeInSeconds = 1
	clientsAmount = 40
	workersAmount = 1000
)

type client struct {
}

type worker struct {
}

type twoForYou struct {
}

type fries struct {
}

type beef struct {
}

type burger struct {
}

func main() {
	start := time.Now()

	clients := make([]client, clientsAmount)
	workersChan := make(chan worker, workersAmount)
	readyOrderChan := make(chan twoForYou, clientsAmount)

	// spawn workers
	for i := 0; i < workersAmount; i++ {
		workersChan <- worker{}
	}

	go func() {
		var wg sync.WaitGroup
		for range clients {
			wg.Add(1)
			go func() {
				orderTwoForYou(workersChan, readyOrderChan)
				wg.Done()
			}()
		}
		wg.Wait()
		close(readyOrderChan)
	}()

	for range readyOrderChan {
		log.Println("Got order")
	}

	log.Printf("Took %s", time.Since(start))
}

func orderTwoForYou(workersChan chan worker, readyOrderChan chan<- twoForYou) {
	stepsM := map[string]func() struct{}{
		"fries":     fryTheFries,
		"beef":      grillTheBeef,
		"burger":    makeBurger,
		"twoForYou": completeTwoForYou,
	}

	var wg sync.WaitGroup
	for step, fun := range stepsM {
		wg.Add(1)
		go func(step string, f func() struct{}) {
			w := <-workersChan
			f()
			wg.Done()
			workersChan <- w
		}(step, fun)
	}
	wg.Wait()

	readyOrderChan <- twoForYou{}
}

//func makeTwoForYou() twoForYou {
//	fryTheFries()
//	_ = grillTheBeef()
//	makeBurger()
//	twoForU := completeTwoForYou()
//	return twoForU
//}

func fryTheFries() struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return struct{}{}
}

func grillTheBeef() struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return struct{}{}
}

func makeBurger() struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return struct{}{}
}

func completeTwoForYou() struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return twoForYou{}
}
