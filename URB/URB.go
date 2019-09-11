package URB

import "fmt"
import "./BEB"
import "./set"
import "strings"

type pending_data struct{
	Module  string
	Message string
}

func Init(address string, addresses []string, input_msgs chan string) chan string {
	delivered 	:= set.New()
	pending 	:= set.New()
	ack			:= make(map[string] *set.Set)


	deliver 	:= make(chan pending_data)


	N 			:= len(addresses)


	sent_msgs := make(chan string)
	rcvd_msgs := BEB.Broadcast(address, addresses, sent_msgs)


	propagate_msgs := make(chan string)


	fmt.Println(delivered.Len())
	fmt.Println(pending.Len())


	go func() { // test DELIVER
		for {
			for _,v := range pending.Elems(){
				pending_message := pending_data{}
				ok := false

				if pending_message, ok = v.(pending_data); !ok {
					continue
				} 

				m := pending_message.Message

				if delivered.Has(m){
					continue
				}

				if 2 * ack[m].Len() <= N{
					continue
				}

				delivered.Insert(m)
				deliver <- pending_message
			}

			select {
				case msg := <- input_msgs:
					f, err 	 := ack[msg]

					if f == nil || err{
						ack[msg] = set.New()
						ack[msg].Insert(address)
					}


					new_pending := pending_data{address, msg}
					pending.Insert(new_pending)

					sent_msgs <- msg


				case rcvd := <- rcvd_msgs:
					tmp := strings.Split(rcvd, "&-&")

					src, msg := tmp[0], tmp[1]

					_, err 	 := ack[msg]

					if err{
						ack[msg] = set.New()
					}

					ack[msg].Insert(src)
					new_pending := pending_data{src, msg}

					if pending.Has(new_pending){
						continue
					}

					pending.Insert(new_pending)
					sent_msgs <- msg

					propagate_msgs <- msg


				default:
					continue
			}
		}
	}()
	return propagate_msgs
}