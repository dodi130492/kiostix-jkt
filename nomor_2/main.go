package main

import "fmt"

func main() {
	for x :=1; x <= 100; x++ {
		if x % 25 == 0 {
			fmt.Println("KI")
		}

		if x % 40 == 0 {
			fmt.Println("OS")
		}

		if x % 60 == 0 {
			fmt.Println("TIX")
		}

		if x % 99 == 0 {
			fmt.Println("KIOSTIX")
		}

	}
}