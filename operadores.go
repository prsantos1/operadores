package main

import "fmt"

func main() {

	for x := 1; x <= 100; x++ {
		if x%3 == 0 {
			fmt.Println("PIN")

		} else if x%5 == 0 {
			fmt.Println("PAN")

		} else if x%3 == 0 && x%5 == 0 {
			fmt.Println("PIN PAN!")

		} else {
			fmt.Println(x)
		}

	}

}
