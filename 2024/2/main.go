package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput() []string {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("failed reading input : %v", err)
	}
	//fmt.Println(data)
	var sb strings.Builder
	length := len(data)
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
func safe_increasing(slice []string) bool {}

func safe_decreasing(slice []string) bool {}

func secondQ() int {
	data := readInput()
	var count int = 0
	for _, x := range data {
		//	fmt.Println(x)
		row := strings.Split(x, " ")
		if safe_increasing(row) || safe_decreasing(row) {
			count += 1
		}
	}

	return count
}

func main() {

	secondQ()

}
