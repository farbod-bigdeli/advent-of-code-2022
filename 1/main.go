package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	res := readData("callories.txt")
	sort.Ints(res)
	fmt.Println("max:", res[len(res)-1])
	fmt.Println("max of top 3:", res[len(res)-1] + res[len(res)-2] + res[len(res)-3])
}

func readData(fname string) ([]int) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	s := bufio.NewScanner(f)
	sumLocal := 0
	res := make([]int, 0)
	for s.Scan() {
		text := s.Text()
		if text == "" {
			res = append(res, sumLocal)
			sumLocal = 0
			
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				log.Fatalf("bad input: %s", err.Error())
			}
			sumLocal += num
		}
	}
	return res
}