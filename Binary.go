package main

import (
	"fmt"
	"strconv"
)

var array = [9]int{256, 128, 64, 32, 16, 8, 4, 2, 1}
var binaryArr [9]int

func main() {
	var value int

	fmt.Print("Enter a value: ")
	fmt.Scanln(&value)

	if value >= 0 {
		decimalToBinary(value)
		fmt.Println()
		decimalToOctet(value)
	}
}

func decimalToBinary(value int) {
	deciInput := value

	// Ask for decimal input
	//fmt.Print("Enter a decimal value: ")
	//fmt.Scanln(&deciInput)

	decimalToBinaryProcess(deciInput)

	// Prints the binary value
	fmt.Print("Binary Value: ")
	for i := 0; i < 9; i++ {
		fmt.Print(binaryArr[i])
	}
}

func decimalToOctet(value int) {
	//deciInput := value
	array2 := [3]int{4, 2, 1}
	var octetArr [3]int
	octetVal := 0

	//decimalToBinaryProcess(deciInput)

	ctr, max := 0, 3

	for i := 0; i < 3; i++ {
		h := 0
		for j := ctr; j < max; j++ {
			if binaryArr[j] == 1 {
				octetVal += array2[h]
				octetArr[i] = octetVal
			}
			h++
		}
		octetVal = 0
		ctr += 3
		max += 3
	}

	// Prints the binary value
	fmt.Print("Octet Value: ")
	for i := 0; i < 3; i++ {
		fmt.Print(octetArr[i])
	}
}

func decimalToBinaryProcess(deciInput int) {
	// Loop for converting the decimal to binary
	for i := 0; i < 9; i++ {
		if deciInput >= array[i] { // If the decimal input is greater than the array value
			deciInput -= array[i] // subtracts the array value to the decimal input
			binaryArr[i] = 1      // and store 1 to the binary array
		} else {
			binaryArr[i] = 0 // Store 0 to the array
		}
	}
}

func binaryToDecimal() {
	var decimal int
	var binaryHolder int
	var binaryArr [100]int
	ctr := -1
	array := [9]int{256, 128, 64, 32, 16, 8, 4, 2, 1}

	// Input a binary and store it in the array
	// Input number "2" to enter the binary inputs
	fmt.Print("Enter a binary digit: ")
	for binaryHolder != 2 {
		ctr++
		fmt.Scan(&binaryHolder)
		binaryArr[ctr] = binaryHolder
	}

	max := 9 // Max size of the array
	if ctr != 9 {
		// Moving the binary digits to the LSB
		for i := ctr; i >= 0; i-- {
			binaryArr[max] = binaryArr[i] // Changing the indext position of the binary
			binaryArr[i] = 0              // Changing the past value of the indext to 0
			max--                         // Decrementing the max for moving the next binary
		}
	}

	// Printing the binary values
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
	fmt.Print("Decimal Value: " + strconv.Itoa(decimal)) // Print the converted binary into decimal
}

func hexToBinary() {
	var tempHolder string
	var hexArr [4]string
	var binary [16]int
	var valueArr [4]int
	arr := [4]int{8, 4, 2, 1}

	// Input a hexadecimal value
	fmt.Print("Enter a hexadecimal: ")
	for i := 0; i < 4; i++ {
		fmt.Scan(&tempHolder)

		// Check if the input hecadecimal is valid
		if tempHolder == "0" ||
			tempHolder == "1" ||
			tempHolder == "2" ||
			tempHolder == "3" ||
			tempHolder == "4" ||
			tempHolder == "5" ||
			tempHolder == "6" ||
			tempHolder == "7" ||
			tempHolder == "8" ||
			tempHolder == "9" ||
			tempHolder == "A" ||
			tempHolder == "B" ||
			tempHolder == "C" ||
			tempHolder == "D" ||
			tempHolder == "E" ||
			tempHolder == "F" {
			hexArr[i] = tempHolder // Store the input hexadecimal to the array
		}
	}

	// Loop for converting each hexademical value into its numerical form or transform it into integer values
	for i := 0; i < 4; i++ {
		if hexArr[i] == "0" {
			valueArr[i] = 0
		} else if hexArr[i] == "1" {
			valueArr[i] = 1
		} else if hexArr[i] == "2" {
			valueArr[i] = 2
		} else if hexArr[i] == "3" {
			valueArr[i] = 3
		} else if hexArr[i] == "4" {
			valueArr[i] = 4
		} else if hexArr[i] == "5" {
			valueArr[i] = 5
		} else if hexArr[i] == "6" {
			valueArr[i] = 6
		} else if hexArr[i] == "7" {
			valueArr[i] = 7
		} else if hexArr[i] == "8" {
			valueArr[i] = 8
		} else if hexArr[i] == "9" {
			valueArr[i] = 9
		} else if hexArr[i] == "A" {
			valueArr[i] = 10
		} else if hexArr[i] == "B" {
			valueArr[i] = 11
		} else if hexArr[i] == "C" {
			valueArr[i] = 12
		} else if hexArr[i] == "D" {
			valueArr[i] = 13
		} else if hexArr[i] == "E" {
			valueArr[i] = 14
		} else if hexArr[i] == "F" {
			valueArr[i] = 15
		}
	}

	// Print the integer values of the hexadecimal
	for i := 0; i < 4; i++ {
		fmt.Print(valueArr[i])
		fmt.Print(" ")
	}

	ctr, max := 0, 4

	// Nested loop for converting the hexadecimal into its binary form
	for i := 0; i < 4; i++ { // Outer loop for locating the index of the hexadecimal values
		value := valueArr[i] // Store the value of a certain index in the "value" variable
		h := 0
		for j := ctr; j < max; j++ { // Inner loop for the convertion
			if value >= arr[h] {
				value -= arr[h]
				binary[j] = 1
			} else if value < arr[h] {
				binary[j] = 0
			}
			h++
		}
		ctr += 4 // Starting point of the next convertion
		max += 4 // Ending point of the convertion
	}

	fmt.Println()

	// Print the converted hexadecimal into binary
	fmt.Print("Binary Value: ")
	for i := 0; i < 16; i++ {
		if i == 4 || i == 8 || i == 12 {
			fmt.Print(" ")
		}
		fmt.Print(binary[i])
	}
}
