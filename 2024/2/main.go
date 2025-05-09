package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("failed reading input : %v", err)
	}
	//fmt.Println(data)
	var sb strings.Builder
	length := len(data) - 1
	for i := 0; i < length; i++ {
		if data[i] <= '9' && '0' <= data[i] {
			sb.WriteByte(data[i])
		} else if data[i] == 10 {
			sb.WriteByte(44)
		} else if data[i] == ' ' {
			sb.WriteByte(data[i])
		}
	}

	s := sb.String()

	return strings.Split(s, ",")
}
func safe_increasing(slice []string) bool {

	for i := 1; i <= len(slice)-1; i++ {
		a, _ := strconv.Atoi(slice[i])
		b, _ := strconv.Atoi(slice[i-1])
		if a <= b {

			return false
		} else if (a - b) > 3 {
			return false
		}

	}
	return true
}

func safe_decreasing(slice []string) bool {

	for i := 1; i <= len(slice)-1; i++ {
		a, _ := strconv.Atoi(slice[i])
		b, _ := strconv.Atoi(slice[i-1])
		if a >= b {
			return false
		} else if (b - a) > 3 {
			return false
		}
	}

	return true
}

func secondQ() int {
	data := readInput()
	var count int = 0
	for _, x := range data {
		row := strings.Split(x, " ")

		if safe_decreasing(row) {
			count += 1
		} else if safe_increasing(row) {
			count += 1
		}
		fmt.Println("\n\n\n\n")
	}
	return count
}

func main() {

	fmt.Println(secondQ())
}
