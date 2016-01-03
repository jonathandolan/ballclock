package clock

//import "fmt"

type Ramp struct {
	contents     []int
	placeholders int
	tippingPoint int
	index        int
	name         string
}

func CreateRamp(size int, placeHolders int, tippingPoint int, name string) Ramp {
	r := Ramp{}
	r.contents = make([]int, size - placeHolders)
	r.placeholders = placeHolders
	r.tippingPoint = tippingPoint - placeHolders
	r.index = 0
	r.name = name
	return r
}

func (r *Ramp) add(input int) (tipped bool, spillOver []int, continued int) {
	if r.index == len(r.contents) {
		tempContents := flip(r.contents)
		r.contents = make([]int, len(r.contents))
		r.index = 0
		return true, tempContents, input
	}
	r.contents[r.index] = input
	r.index++
	return false, []int{0}, 0
}

func flip(input []int) []int {
	length := len(input) - 1
	temp := make([]int, length+1)
	for i := 0; i < length+1; i++ {
		temp[length-i] = input[i]
	}
	return temp
}
