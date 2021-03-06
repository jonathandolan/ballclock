package clock

import (
	"fmt"
	"os"
)

type Clock struct {
	queue              []int
	size               int
	index              int
	minutes            Ramp
	fiveMinutes        Ramp
	hours              Ramp
	minutesElapsed     int
}

func CreateClock(size int) Clock {
	c := Clock{}
	c.queue = make([]int, size)
	c.size = size
	c.minutes = CreateRamp(4, 0, 5, "Minutes")
	c.fiveMinutes = CreateRamp(11, 0, 12, "Five minutes")
	c.hours = CreateRamp(11, 1, 12, "Hours")
	c.index = 0

	for i := 0; i < size; i++ {
		c.queue[i] = i + 1
	}

	return c
}

func (c *Clock) Process() int {
	for {
		c.addOneMinute()
		if c.isQueueRepeated(){
			break
		}
	}
	fmt.Print("size: ")
	fmt.Println(c.size)
	fmt.Print("minutes: ")
	fmt.Println(c.minutesElapsed)

	return 7
}

func (c *Clock) PrintStateAtIteration(iteration int){
	for i := 0; i < iteration + 1; i++{
		c.addOneMinute()
	}

	fmt.Println()
	fmt.Print("queue: ")
	fmt.Println(c.queue)
	fmt.Print("minutes contents: ")
	fmt.Println(c.minutes.contents)
	fmt.Print("five minutes contents: ")
	fmt.Println(c.fiveMinutes.contents)
	fmt.Print("hours contents: ")
	fmt.Println(c.hours.contents)
	fmt.Print("size: ")
	fmt.Println(c.size)
	fmt.Print("is queue repeated: ")
	fmt.Println(c.isQueueRepeated())
}

func (c *Clock) addOneMinute() {
	tipped, spill, continued := c.minutes.add(c.GetFromQueue())
	c.updateTimeAddMinute()
	c.index++
	if tipped {
		c.AddToQueue(spill)
		c.addFiveMinutes(continued)
	}
}

func (c *Clock) addFiveMinutes(input int) {
	tipped, spill, continued := c.fiveMinutes.add(input)
	if tipped {
		c.AddToQueue(spill)
		c.addOneHour(continued)
	}

}

func (c *Clock) addOneHour(input int) {
	tipped, spill, continued := c.hours.add(input)
	if tipped {
		c.AddToQueue(spill)
		c.AddToQueue([]int{continued})
	}

}

func (c *Clock) AddToQueue(input []int) {
	queueCopy := c.queue
	for i := 0; i < len(input); i++ {
		queueCopy = append(queueCopy, input[i])
	}
	c.queue = queueCopy
}

func (c *Clock) GetFromQueue() int {
	if len(c.queue) == 0 {
		return 0
	}

	answer := c.queue[0]
	tempQueue := c.queue[1:]
	c.queue = tempQueue

	return answer
}

func (c *Clock) updateTimeAddMinute() {
	c.minutesElapsed++
}

func (c *Clock) TestQueueRepeated(){
	for {
		fmt.Println()
		fmt.Print("queue: ")
		fmt.Println(c.queue)
		fmt.Print("is queue repeated: ")
		fmt.Println(c.isQueueRepeated())
		fmt.Print("minutes contents: ")
		fmt.Println(c.minutes.contents)
		fmt.Print("five minutes contents: ")
		fmt.Println(c.fiveMinutes.contents)
		fmt.Print("hours contents: ")
		fmt.Println(c.hours.contents)
		fmt.Print("size: ")
		fmt.Println(c.size)
		fmt.Println()
		c.addOneMinute()


		if c.isQueueRepeated(){
			fmt.Println(c.queue)
			os.Exit(1)
		}

	}

}



func (c *Clock) isQueueRepeated() bool{
	if len(c.queue) == c.size {
		for i := 0; i < len(c.queue); i++ {
			if i + 1 != c.queue[i] {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

