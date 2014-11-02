package main

import (
	_ "betofahl"
	"fmt"
	"time" 
)

func main(){
	fmt.Println("Ïch laufe")
	<-time.NewTimer(30*time.Second).C
}