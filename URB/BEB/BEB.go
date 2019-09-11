package BEB

import . "../Link"
// import "fmt"


func Broadcast(address string, addresses []string, msgs chan string) chan string{
	link := PP2PLink{}
	link.Ind = make(chan PP2PLink_Ind_Message)
	link.Req = make(chan PP2PLink_Req_Message)

	link.Init(address)
	rcvd := make(chan string)

	go func(){
		for {
			msg := <- msgs
			// go func (){
			for _, addr := range addresses {
				if addr == address{
					continue
				}

				M := PP2PLink_Req_Message{addr, address+"&-&"+msg}
				link.Req <- M
				// fmt.Println(M)
			}
			// }()
		}
	}()

	go func(){
		for {
			m:=<-link.Ind
			// fmt.Println(m)
			// src, msg := s[0], s[1]
			// src = src
			
			// rcvd <- msg
			rcvd <- m.Message
		}
	}()

	return rcvd
}