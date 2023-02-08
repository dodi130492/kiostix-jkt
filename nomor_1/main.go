package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var kecil, besar int
		for i, e := range n {
			switch {
			case i == 0:
				besar, kecil = e, e
			case e > besar:
				besar = e
			case e < kecil:
				kecil = e
			}
		}
		return kecil, besar
	}
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var kecil, besar = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, kecil, besar)
}