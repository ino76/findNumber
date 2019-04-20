/*
	This program should iterate through numbers and try to multiplicate
    its digits between themselves while counting how many multiplications
    is possible until there is only one digit
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Function for finding a number ...
func findNumber() {

	findMultiplicationOf, _ := strconv.Atoi(os.Args[1])
	var number uint64 = 0
	maxMulti := 0

	for number < 99999999999999 {

		currentNumber := number
		currentMulti := 0

		for getLen(currentNumber) > 1 {
			currentNumber = multiplyDigits(currentNumber)
			currentMulti++
		}

		if currentMulti > maxMulti {
			maxMulti = currentMulti
		}

		// if we found multi high like we specified log message and break a loop
		if maxMulti == findMultiplicationOf {
			fmt.Printf("\nNumber %d has multiplication persistance of %d.\n\n", number, maxMulti)
			printMulti(number)
			break
		}

		number++
	}

}

func printMulti(number uint64) {
	fmt.Println("------------------------")
	fmt.Println("n: ", number)
	fmt.Println("------------------------")
	currentMulti := 1
	for getLen(number) > 1 {
		number = multiplyDigits(number)
		fmt.Printf("%d.  %d\n", currentMulti, number)
		currentMulti += 1
	}
	fmt.Println("________________________")
}

// function that multiply digits of a number and return result
func multiplyDigits(number uint64) uint64 {
	mult := uint64(1)

	for number != 0 {
		mult *= number % 10
		number /= 10
	}

	return mult
}

// function that return length of a number
func getLen(number uint64) (count uint64) {
	for number != 0 {
		number /= 10
		count++
	}
	return
}

func main() {
	start := time.Now()
	findNumber()
	fmt.Printf("Finding this number took %s", time.Since(start))
}
