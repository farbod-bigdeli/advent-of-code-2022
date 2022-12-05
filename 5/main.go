package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data := readData("input1")
	for _, v := range data {
		v.Reverse()
	}
	proccessActions("input2", data)
	for i, v := range data {
		s, _ := v.Pop()
		fmt.Println(i, s)
	}
}

func readData(fname string) (map[int]Stack) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	res := make(map[int]Stack)
	s := bufio.NewScanner(f)
	for s.Scan() {
		text := s.Text()
		for i := 0; i < len(text); i++ {
			if string(text[i]) == "[" {
				key := i / 4 + 1
				s, ok := res[key]
				if ok {
					s.Push(string(text[i + 1]))
					res[key] = s
				} else {
					s := Stack{}
					s.Push(string(text[i + 1]))
					res[key] = s
				}
			}
		}
	}
	return res
}

func proccessActions(fname string, in map[int]Stack) (map[int]Stack) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	res := make(map[int]Stack)
	s := bufio.NewScanner(f)
	for s.Scan() {
		s := strings.NewReader(s.Text())
		var source, dest, count int
		fmt.Fscanf(s, "move %d from %d to %d", &count, &source, &dest)
		moveCrate(source, dest, count, in)
	}
	return res
}

func moveCrate(source, dest, count int,in map[int]Stack) {
	s1 := in[source]
	s2 := in[dest]
	s3 := Stack{}
	for i := 0; i < count; i++ {
		str, _ := s1.Pop()
		s3.Push(str)	
	}
	
	for i := 0; i < count; i++ {
		str, _ := s3.Pop()
		s2.Push(str)
	}

	in[source] = s1
	in[dest] = s2
}