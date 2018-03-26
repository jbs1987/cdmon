package main

import (
	"fmt"
)

//function that determines if two numbers are multiples
func multiplo(num, mult int) bool {

	if num%mult == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	//Initialize the counter in 1
	num := 1

	for num <= 100 { //counter ends in 100
		//if the number of countes is multiple of 3 print CD
		if multiplo(num, 3) == true && multiplo(num, 5) == false {
			fmt.Println("\x1b[33;1mCD\x1b[0m")
		}
		//if the number of countes is multiple of 5 print CD
		if multiplo(num, 3) == false && multiplo(num, 5) == true {
			fmt.Println("\x1b[34;1mmon\x1b[0m")
		}
		//if the number of countes is multiple of 3 and 5 at same time, print CDmon
		if multiplo(num, 3) == true && multiplo(num, 5) == true {
			fmt.Println("\x1b[32;1mCDmon\x1b[0m")
		}
		//rest of cases
		if multiplo(num, 3) == false && multiplo(num, 5) == false {
			fmt.Println(num)
		}
		num++
	}
}
