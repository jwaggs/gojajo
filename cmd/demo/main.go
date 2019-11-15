package main
import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

// a wrapper to always print when we call wg.Done
func decrementWg() {
	// wg.Done() subtracts one from the waitgroup. we explicitly added 2 to the wait group as the first step in main.
	wg.Done()
	fmt.Println("removed one from the wait group")
}

// simple function to wait for the time specified before returning.
func waitFor(t time.Duration) {
	fmt.Println("entering waitFor with a value of", t)
	// defer is a way to declare something must execute after the function returns no matter what.
	// defer is helpful when you want to make sure something runs, even if an error occurs.
	// defer our custom decrement function
	defer decrementWg()

	// this will sleep the goroutine for the time specified.
	time.Sleep(t)
	// this line will then be executed after time.Sleep returns
	fmt.Println("exiting waitFor with a value of", t)
	// you'll notice that the print statement defined in decrementWg actually occurs here.
}

func main() {
	fmt.Println("entering main")
	// add 2 to the waitgroup... this means that two wg.Done() calls will need to be made before a wg.Wait() call will return.
	// see below for why I did this.
	wg.Add(2)
	// fmt.Println("added two to the wait group")
	// go is the keyword to start a new goroutine...
	// this invokes the "waitFor" function on a new, separate go routine concurrently, while this (main) function continues on without waiting for the result.
	go waitFor(10 * time.Second)
	// again, this starts another go routine.
	go waitFor(11 * time.Second)
	// we now have 3 separate go routines running in parallel.. this one, the 1 second wait for call and the 5 second waitFor call.
	// so how do we stop the program from exiting before we are all done?
	// wg.Wait will not return until the waitgroup delta returns to zero.
	fmt.Println("WAITING")
	// wg.Wait()
	fmt.Println("exiting program")
}
