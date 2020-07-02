/**
 * go channels implementation practice
 */
package main

const N = 20

/**
 * input numbers into an one-direction channel
 */
func producer(out chan<- int) {
	for i := 0; i < N; i++ {
		out <- i
	}
}

/**
 * receive numbers from input channel, process them with adder and send to output channel
 */
func adder(in <-chan int, out chan<- int, val int) {
	for {
		x := <-in
		out <- x + val
	}
}

/**
 * receive numbers from input channel, process them with multiplier and send to output channel
 */
func multiplier(in <-chan int, out chan<- int, val int) {
	for {
		x := <-in
		out <- x * val
	}
}

/**
 * receive numbers from input channel, print the numbers, output "true" to done channel
 */
func consumer(in <-chan int, done chan<- bool, val int) {
	for i := 0; i < N; i++ {
		x := <-in
		if i%val == 0 {
			print(x, "\n")
		}
	}
	done <- true
}

func main() {
	addVal := 3
	multVal := 2
	filterVal := 3
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	done := make(chan bool)
	go producer(ch1)
	go adder(ch1, ch2, addVal)
	go multiplier(ch2, ch3, multVal)
	go consumer(ch3, done, filterVal)
	<-done
}
