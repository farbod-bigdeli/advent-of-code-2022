package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


type input struct {
	left, right string
}
func main()  {
	inp := readData("input")
	c := 0
	// for i := range inp {
	// 	if inp[i].hasFullOverlap() {
	// 		c++
	// 	}
	// }
	// fmt.Println(c)
	for i := range inp {
		if inp[i].hasOverlap() {
			c++
		}
	}
	fmt.Println(c)
}
func readData(fname string) ([]input) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	s := bufio.NewScanner(f)
	res := make([]input, 0)
	for s.Scan() {
		text := s.Text()
		inpArr := strings.Split(text, ",")
		rps := input{inpArr[0], inpArr[1]}
		res = append(res, rps)
	}
	return res
}

func (in input) hasFullOverlap() bool {
	arr := strings.Split(in.left, "-")
	minL, _ := strconv.Atoi(arr[0])
	maxL, _ := strconv.Atoi(arr[1])
	arr = strings.Split(in.right, "-")
	minR, _ := strconv.Atoi(arr[0])
	maxR, _ := strconv.Atoi(arr[1])
	if (minL <= minR && maxL >= maxR) || (minR <= minL && maxR >= maxL) {
		return true
	}
	return false


}

func (in input) hasOverlap() bool {
	arr := strings.Split(in.left, "-")
	minL, _ := strconv.Atoi(arr[0])
	maxL, _ := strconv.Atoi(arr[1])
	arr = strings.Split(in.right, "-")
	minR, _ := strconv.Atoi(arr[0])
	maxR, _ := strconv.Atoi(arr[1])
	if (minL <= minR && minR <= maxL) || (maxR <= maxL && maxR >= minL) ||
	(minR <= minL && minL <= maxR) || (maxL <= maxR && maxL >= minR) {
		return true
	}
	return false


}