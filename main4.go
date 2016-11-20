package main

import (
  "fmt"
)

func main() {
  isAlive := false
  
  for !isAlive {
	  
	  if  PingChatServer() {
         isAlive = true
       }
  }

  fmt.Println("Ping is done")
}

func PingChatServer()bool{
	fmt.Println("buradaydım")
	return true
}