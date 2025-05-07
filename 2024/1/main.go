package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Printf("%T", data)

	// making two seperate slices of thease two column of code
	var sb strings.Builder
	var count int8 = 0
	// first deleteing \n and replacing it with space
	for i := 0; i < len(data); i++ {
		if data[i] <= '9' && '0' <= data[i] {
			sb.WriteByte(data[i])
		}
		if data[i] == 10 {
			sb.WriteByte(44)
		}
		if data[i] == ' ' {
			count += 1
			if count == 3 {
				sb.WriteByte(44)
				count = 0
			}
		}

	}
	var s string
	s = sb.String()
	//fmt.Println(s)

	slice := strings.Split(s, ",")
	fmt.Println(slice)
	var Right, Left []int

	for i := 0; i < len(slice); i++ {
		p, _ := strconv.Atoi(slice[i])
		if i%2 != 0 {

			Left = append(Left, p)
		} else {
			Right = append(Right, p)
		}
	}

	sortSlice := func(slice *[]int) {

	}

	sortSlice(&Left)
	sortSlice(&Right)
}
