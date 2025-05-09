package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func merge(slice *[]int, l int, m int, h int) {

	var L, H []int
	n1 := m - l + 1
	n2 := h - m
	for i := 0; i < n1; i++ {
		L = append(L, (*slice)[l+i])

	}
	for j := 0; j < n2; j++ {
		H = append(H, (*slice)[m+1+j])
	}

	var (
		i = 0
		j = 0
		k = l
	)
	for i < n1 && j < n2 {
		if L[i] <= H[j] {
			(*slice)[k] = L[i]
			i++
		} else {
			(*slice)[k] = H[j]
			j++
		}
		k++

	}

	for i < n1 {
		(*slice)[k] = L[i]
		i++
		k++

	}

	for j < n2 {
		(*slice)[k] = H[j]
		j++
		k++

	}
}

func mergeSort(slice *[]int, l int, h int) {

	if l < h {
		m := l + (h-l)/2
		mergeSort(slice, l, m)
		mergeSort(slice, m+1, h)

		merge(slice, l, m, h)

	}

}

func firstQ() int {

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
	//fmt.Println(slice)
	var Right, Left []int

	for i := 0; i < len(slice)-1; i++ {
		p, _ := strconv.Atoi(slice[i])
		if i%2 != 0 {

			Left = append(Left, p)
		} else {
			Right = append(Right, p)
		}
	}
	length := len(Left)
	mergeSort(&Left, 0, length-1)
	mergeSort(&Right, 0, length-1)
	/*
		fmt.Printf("%v and the length :%d", Left, len(Left))
		fmt.Println("\n\n\n")
		fmt.Printf("%v and the length :%d", Right, len(Right))
	*/
	var sum int
	for i := 0; i < length; i++ {
		num := Left[i] - Right[i]
		if num < 0 {
			num *= -1
		}
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(firstQ())
}
