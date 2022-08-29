package main

import (
	"log"
	"time"
)

const timeInSeconds = 1

func main() {
	start := time.Now()
	_ = makeTwoForYou()
	log.Printf("Took %s", time.Since(start))
}

func makeTwoForYou() struct{} {
	fries := fryTheFries()
	beef := grillTheBeef()
	burger := makeBurger(beef)
	twoForYou := completeTwoForYou(burger, fries)
	return twoForYou
}

func fryTheFries() struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return struct{}{}
}

func grillTheBeef() struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return struct{}{}
}

func makeBurger(beef struct{}) struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return struct{}{}
}

func completeTwoForYou(burger struct{}, fries struct{}) struct{} {
	time.Sleep(timeInSeconds * time.Second)
	return struct{}{}
}
