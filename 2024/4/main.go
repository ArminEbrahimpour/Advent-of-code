package main

import (
	"fmt"
	"os"
)

func ReadData() int {
	data, _ := os.ReadFile("./input.txt")
	var count int = 0
	length := len(data)
	for i := 0; i < length-4; i++ {
		// horizental
		if (data[i] == 'X' && data[i+1] == 'M') && (data[i+2] == 'A' && data[i+3] == 'S') {
			count += 1
		}
		if (data[i] == 'S' && data[i+1] == 'A') && (data[i+2] == 'M' && data[i+3] == 'X') {
			count += 1
		}
		// vertical (maximum lenght of horizenal length is 140 so for being vertically below it means we can just add 140 to the i)
		v := i + (3 * 140)
		if v <= length {
			if (data[i] == 'S' && data[i+140] == 'A') && (data[i+(2*140)] == 'M' && data[i+(3*140)] == 'X') {
				count += 1
			}
			if (data[i] == 'X' && data[i+140] == 'M') && (data[i+(2*140)] == 'A' && data[i+(3*140)] == 'S') {
				count += 1
			}
		}
		// now we need to add the diagonal

		//right diagonal
		if v+3 <= length {
			if (data[i] == 'S' && data[i+140+1] == 'A') && (data[i+(2*140)+2] == 'M' && data[i+(3*140)+3] == 'X') {
				count += 1
			}
			if (data[i] == 'X' && data[i+140+1] == 'M') && (data[i+(2*140)+2] == 'A' && data[i+(3*140)+3] == 'S') {
				count += 1
			}
		}
		//left diagonal
		if v <= length {
			if (data[i] == 'S' && data[i+140-1] == 'A') && (data[i+(2*140)-2] == 'M' && data[i+(3*140)-3] == 'X') {
				count += 1
			}
			if (data[i] == 'X' && data[i+140-1] == 'M') && (data[i+(2*140)-2] == 'A' && data[i+(3*140)-3] == 'S') {
				count += 1
			}
		}
	}

	return count
}

func main() {
	fmt.Println(ReadData())
}
