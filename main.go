package main

import (
	"./clock"
	"fmt"
)

func main() {
	fmt.Println("Ball clock")
	c := clock.CreateClock(30)
//	c.PrintStateAtIteration(325)
	c.Process()
//	c.TestQueueRepeated()
}




//
//{"Min":[],"FiveMin":[22,13,25,3,7],"Hour":[6,12,17,4,15],"Main"
//[11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9]}