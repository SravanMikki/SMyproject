// syntax: - mychannel := make(ch chan <-int)
// sending the type into channel
// mychannel <- 20
// receving the type channel
// result:= <- mychannel

package main

import "fmt"

func Sendmessage(ch chan<- int) {
	ch <- 23
}

func main() {
	ch := make(chan int)
	go Sendmessage(ch)
	message := <-ch
	fmt.Println("Message recived: ", message)
}
