package main


import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"strconv"
)

func main() {

	d := make(chan os.Signal, 1)
	signal.Notify(d, os.Interrupt)

	go func() {
		<-d
		fmt.Println("Operation stopped. Signal interrupted")
		os.Exit(1)
	}()

	c := make(chan int)
	go readInput(c)
	time.Sleep(5 * 1e9)
	go addUp(c)
	time.Sleep(5 * 1e9)
}

func readInput(c chan int) {

	var n1 string
	var n2 string

	fmt.Println("Enter num: ")
	fmt.Scan(&n1)

	tall1, err := strconv.Atoi(n1)
	if (err != nil) {
		fmt.Println("Input not valid. Needs a type int")
		os.Exit(0)
	}

	fmt.Println("Enter num: ")
	fmt.Scan(&n2)

	tall2, err := strconv.Atoi(n2)
	if (err != nil) {
		fmt.Println("Input not valid. Needs a type int")
		os.Exit(0)
	}


	c <- tall1 //sender data via channel
	c <- tall2

	res := <-c // mottar resultat fra channel
	fmt.Println("Result: ", res)

}

func addUp(c chan int) {

	n1, n2 := <-c, <-c // mottar data fra readInput()
	res := (n1 + n2)

	c <- res // sender resultat tilbake til readInput()

}
