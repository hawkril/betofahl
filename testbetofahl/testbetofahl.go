package main

import (
	"github.com/hawkril/betofahl/betofahl"
	"fmt"
	"time" 
)

func main(){
	defer betofahl.Defer()

	fmt.Println("Ïch laufe")
	<-time.NewTimer(30*time.Second).C
}