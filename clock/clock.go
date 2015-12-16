package clock

import ("fmt")

type Clock struct{
	queue []int
	size int
	index int
	minutes Ramp
	fiveMinutes Ramp
	hours Ramp
}


func CreateClock(size int) Clock{
	c := Clock{}
	c.queue = make([]int, size)
	c.size = size
	c.minutes = CreateRamp(4, 0, 5)
	c.fiveMinutes = CreateRamp(11, 0, 12)
	c.hours = CreateRamp(11, 1, 12)
	c.index = 0
	return c
}

func (c *Clock) Process(numberOfBalls int) int {
	for i := 0; i < numberOfBalls; i++{
		c.queue[i] = i + 1
	}

	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)
	c.addOneMinute(numberOfBalls)

	return numberOfBalls
}

func (c *Clock) addOneMinute(numberOfBalls int){

	if c.index == numberOfBalls - 1{
		c.index = 0
	}
	tipped, minSpill, minCont := c.minutes.add(c.queue[c.index + 1])
	c.index++

	if tipped{
		fmt.Print("minSpill: ")
		fmt.Println(minSpill)
		fmt.Print("minCont: ")
		fmt.Println(minCont)
	}
}
