package main

import (
  "fmt"
)

func main() {
 
  for {
	  if PingChatServer() {
        break
       }
  }
  fmt.Println("Ping is done")
}

func PingChatServer()bool{
	fmt.Println("buradaydım ben")
	return true
}