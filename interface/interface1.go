package main

import "fmt"

type spieces interface {
	skill() string
	running() string
}
type humans struct {
	region int
	sName  string
}
type dogs struct {
}

// Implement `skill()` method for `humans`
func (h humans) skill() {
	fmt.Println(h.region)
	fmt.Println(h.sName)

}

func (d dogs) running() string {
	return "fast"

}

func (d dogs) skill() string {
	return "Dogs have strong senses"
}

func All(s spieces) {
	fmt.Println(s.running())
}

func main() {
	h1 := humans{91, "homa"}
	h1.skill()

	d1 := dogs{}
	fmt.Println(d1.skill())
	All(d1)

}
