package main

import (
		"fmt"
		"time"
		)

func main(){
	c1 :=make(chan string)
	c2 :=make(chan string)
	//c := make(chan int, 1)

		go func(){
			for{
				c1 <- "from 1"
				time.Sleep(time.Second*1)

			}
		}()
		go func(){
			for{
				c2 <- "from 2"
				time.Sleep(time.Second*2)

			}
		}()
		go func(){
			for i:=0; ; i++{
				select{
				case msg1 := <-c1:
					fmt.Println(msg1)
				case msg2 := <-c2:
					fmt.Println(msg2)
				//case <- time.Tick(time.Second):
				//	fmt.Println(statusUpdate())
				case <-time.After(time.Second):
					fmt.Println("timeout")
				//default:
				//	fmt.Println("nothing ready")
				}
			}
		}()
		var input string
		fmt.Scanln(&input)

}