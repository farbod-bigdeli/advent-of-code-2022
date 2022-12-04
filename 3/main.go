package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type input struct {
	left, right string
}

type input2 struct {
	first, second, third string
}

func main()  {
	// inp := []input{}
	// readData("input", &inp)
	// sum := 0
	// for i := range inp {
	// 	calcDuplicateChar(inp[i].left, inp[i].right, &sum)
	// }
	// fmt.Println(sum)

	inp := []input2{}
	readDataV2("input", &inp)
	sum := 0
	for i := range inp {
		calcDuplicateCharV2(inp[i].first, inp[i].second,inp[i].third, &sum)
	}
	fmt.Println(sum)
}

func readData(fname string, in *[]input) () {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		text := s.Text()
		mid := len(text)/2
		left := text[:mid]
		right := text[mid:]
		*in = append(*in, input{left , right})
	}
}

func readDataV2(fname string, in *[]input2) () {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	s := bufio.NewScanner(f)
	index := 0
	inputs := make([]string, 3)
	for s.Scan() {
		text := s.Text()
		inputs[index] = text
		index++
		if index == 3 {
			*in = append(*in, input2{
				first: inputs[0],
				second: inputs[1],
				third: inputs[2],
			})
			index = 0
		}
	}
}
func calcDuplicateChar(s1, s2 string, sum *int) {
	checked := make(map[string]bool, 0)
	for i := range s1 {
		for j := range s2{
			if s1[i] == s2[j] {
				if _, ok := checked[string(s1[1])]; !ok {
					checked[string(s1[1])] = true
					if (int(s1[i]) > 96) &&  (int(s1[i]) < 123){
						*sum += int(s1[i]) - 96
					} else {
						*sum += int(s1[i]) - 38
					}
				}
				
			}
		}
	}
}

func calcDuplicateCharV2(s1, s2, s3 string, sum *int) {
	checked := make(map[string]bool, 0)
	for i := range s1 {
		for j := range s2{
			if s1[i] == s2[j] {
				for k := range s3 {
					if s3[k] == s1[i] {
						if _, ok := checked[string(s1[i])]; !ok {
							fmt.Println(string(s1[i]))
							checked[string(s1[i])] = true
							if (int(s1[i]) > 96) &&  (int(s1[i]) < 123){
								*sum += int(s1[i]) - 96
							} else {
								*sum += int(s1[i]) - 38
							}
						}

					}
				}
					
				
				
			}
		}
	}
}