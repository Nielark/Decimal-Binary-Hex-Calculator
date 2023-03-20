package main

import "fmt"

func main() {
	fmt.Println("DECIMAL & BINARY & HEXADECIMAL CONVERTER")

	//hexToBinary()
}

func binaryToDecimal() {
	var decimal int
	var binaryHolder int
	var binaryArr [100]int
	ctr := -1
	array := [9]int{256, 128, 64, 32, 16, 8, 4, 2, 1}

	// Input a binary and store it in the array
	fmt.Print("Enter a binary digit: ")
	for binaryHolder != 2 {
		ctr++
		fmt.Scan(&binaryHolder)
		binaryArr[ctr] = binaryHolder
	}

	max := 9 // Max size if the array
	if ctr != 9 {
		for i := ctr; i >= 0; i-- {
			binaryArr[max] = binaryArr[i]
			binaryArr[i] = 0
			max--
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Print(binaryArr[i])
	}

	// Convert the binary into decimal
	for i := 0; i < 9; i++ {
		if binaryArr[i] == 1 { // Check if the bit is equal to one
			decimal += array[i] // Add the values in the array and stored it in the decimal if the the condition is true
		}
	}
	fmt.Println()
	fmt.Print(decimal) // Print the decimal value of the binary
}

func hexToBinary() {
	//var decimal int
	var hexArr [4]string
	var tempHolder string

	fmt.Print("Enter a hexadecimal: ")
	for i := 0; i < 4; i++ {
		fmt.Scan(&tempHolder)

		if tempHolder == "1" ||
			tempHolder == "2" ||
			tempHolder == "3" ||
			tempHolder == "4" ||
			tempHolder == "5" ||
			tempHolder == "6" ||
			tempHolder == "7" ||
			tempHolder == "8" ||
			tempHolder == "9" ||
			tempHolder == "0" ||
			tempHolder == "A" ||
			tempHolder == "B" ||
			tempHolder == "C" ||
			tempHolder == "D" ||
			tempHolder == "E" ||
			tempHolder == "F" {
			hexArr[i] = tempHolder
		}
	}

	for i := 0; i < 4; i++ {
		if hexArr[i] == "1" {

		}
	}
}
