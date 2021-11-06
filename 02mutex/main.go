package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	/*IIFEE*/

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One Routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three Routine")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)

	/*
		Run go run --race main.go
	*/
	/*
		They are not execting in order and will not
	*/

	/*
		RWMutex concept:
		rwmut ko use karte hai lumutex ki jagah jab
		koi goroutine ek meory ko read karta hai to vo use
		read karne data hai agar koi aur goroutine read karta hai
		to bhi read karne deta hai multiple go routines are allowed
		to read a memory lekin jab ek go routine ko write karni ho
		memory to jitne bhi reader hote hai unko vo kick karta hai
		aur tab tak read nahi karne deta jabtak write karne vala gogorutine
		apna exectuin khtm  na karde.
	*/
}
