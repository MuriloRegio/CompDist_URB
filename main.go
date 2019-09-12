package main

import "fmt"
import "os"
import "./URB"
import "strconv"
import "math/rand"
import "time"
	
var max = 10
var s2  = rand.NewSource(42)
var r2  = rand.New(s2)

func body (channel chan string) {
	fail := r2.Intn(max)
	die  := r2.Intn(N) == id

	fmt.Println(fail)
	fmt.Println(die)

	for i := 0; i < max; i++ {
		msg := strconv.Itoa(i)

		if i == fail && die{
			msg = "fail"
		}

		msg = msg+"-"+address
		channel <- msg

		if i == fail && die{
			time.Sleep(5 * time.Millisecond)
			a := 1
			a  = a / (1 - a)
		}
	}
}

var N = 4
var addresses = []string{}
var address = ""
var id = 0

func mkConnections() {
	for i := 0; i < N; i++ {
		addresses = append(addresses, "127.0.0.1:500"+strconv.Itoa(i))
	}
}

func main (){
	args  := os.Args[1:]

	if len(args) > 1 {
		index, err := strconv.Atoi(args[1])

		if err != nil{
			fmt.Println("Invalid parameter:\t"+args[1])
			return
		}
		N = index
	}

	mkConnections()

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
	fmt.Println(address)

	// for i := 0; i <id; i++ {
	// 	_ = r2.Intn(max)
	// }

	input := make(chan string)

	rcvd  := URB.Init(address, addresses, input)

	time.Sleep(3*time.Second)

	go body(input)

	go func() {
		for { 
			select {
				case msg := <-rcvd:
					fmt.Println("->"+msg)
				default:
					continue
			}
		}
	}()

	
	time.Sleep(3*time.Second)
}