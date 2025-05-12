package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readData() string {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("reading data failed: %v", err)
	}

	//fmt.Printf("%T", data)
	return string(data)
}

func main() {
	data := readData()
	//fmt.Println(data)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	muls := re.FindAllString(data, -1)
	//fmt.Println(muls)
	var sum int
	for _, x := range muls {
		fmt.Println(x)

		var sb strings.Builder
		var sa strings.Builder
		flag := true
		for i := 0; i < len(x); i++ {
			if x[i] <= '9' && x[i] >= '0' && flag == true {
				sb.WriteByte(x[i])
			}
			if x[i] == ',' {
				flag = false
			}
			if x[i] <= '9' && x[i] >= '0' && flag == false {
				sa.WriteByte(x[i])
			}
		}
		a, _ := strconv.Atoi(sa.String())
		b, _ := strconv.Atoi(sb.String())

		sum += (a * b)

	}
	fmt.Println(sum)
}
