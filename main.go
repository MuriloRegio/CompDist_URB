package main

import "fmt"
import "os"
import "./URB"
import "strconv"
import "math/rand"
	
var max = 1000
var s2  = rand.NewSource(42)
var r2  = rand.New(s2)

func body (channel chan string) {
	fail := r2.Intn(max)

	for i := 0; i < max; i++ {
		msg := strconv.Itoa(i)

		if i == fail {
			msg = "fail"
		}

		msg = msg+"-"+address
		channel <- msg
	}
}

var N = 10
var addresses = []string{}
var address = ""

func mkConnections() {
	for i := 0; i < N; i++ {
		addresses = append(addresses, "127.0.0.1:500"+strconv.Itoa(i))
	}
}

func main (){
	mkConnections()

	args  := os.Args[1:]
	id 	  := 0

	if len(args) > 0 {
		index, err := strconv.Atoi(args[0])

		if err != nil{
			fmt.Println("Invalid parameter:\t"+args[0])
			return
		}
		id = index
	}else{
		id = 0
	}

	address = addresses[id]

	for i := 0; i <id; i++ {
		_ = r2.Intn(max)
	}

	input := make(chan string)

	rcvd  := URB.Init(address, addresses, input)

	go body(input)

	for { 
		select {
			case msg := <-rcvd:
				fmt.Println(msg)
			default:
				continue
		}
	}
}