package main

import ("fmt"
		"./clock"
)

func main(){
	fmt.Println("hello worlds")
	c := clock.CreateClock(27)
	c.Process(27)
//	c.AddToRamp()
}
