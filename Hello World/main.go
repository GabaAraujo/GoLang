package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {

		os.Exit(1)
	}

	fmt.Println("Hello", os.Args[1])

	//fmt.Println("Hello World")

	/*
	   rateLimit := time.Tick(500 * time.Millisecond)
	   <-rateLimit

	   var wg sync.WaitGroup
	   for i := 0; i < 10; i++ {
	       wg.Add(1)
	       go func(i int) {
	           <-rateLimit
	           fmt.Println("Hello", i)
	           wg.Done()
	       }(i)
	   }
	   wg.Wait()
	*/

}
