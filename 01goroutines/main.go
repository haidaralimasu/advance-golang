package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	/*
		go greeter("Hello")
		greeter("World")
	*/
	websitelist := []string{
		"https://blog.kodinghandle.com",
		"https://haidar.vercel.app",
		"https://github.com",
		"https://bullcarter.in",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // At end every time
	fmt.Println(signals)
	/*
		go goroutine threads ko use karta hai
		lekin agar ek se jyada go goroutine ek
		hi memory ke alloct kare tab go routines
		slow ho jate hai reson dene ki jarurat nahi
		honi chaiye aur ye acha bhi nahi hai

		isliye hum use karte hai mutex
		mutex urf mutual execlution lock yani ki
		ek goroutine ko ek particular memory allot ki jati hai
		aur jitne b goroutine run ho rahe ho unko utni memory allot
		ki jati hai taki sab goroutine ek sath execute ho sake parellism
	*/

	/*
		lock unlock mutex
		jab tak ek goroutine run ho rha hai tab tak ek memory
		allot kar di jayegi aur tab tak kuch aur use use nahi kar payega
	*/

}

/*
func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
*/

func getStatusCode(uri string) {
	defer wg.Done()

	res, err := http.Get(uri)

	if err != nil {
		fmt.Println("Invalid uri")
	}

	mut.Lock()
	signals = append(signals, uri)
	mut.Unlock()

	fmt.Printf("%d status code for %s  \n", res.StatusCode, uri)
}
