package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main()  {
	proccessStream("input")
}

func proccessStream(fname string) () {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	data, _ := io.ReadAll(f)
	for i := range data {
		b := make([]byte, 14)
		n, _ := f.ReadAt(b, int64(i))
		if n == 14 && !hasDuplicate(b) {
			fmt.Println(i + 14)
			return
		}
	
	}
}
func hasDuplicate(b []byte) bool {
	m := make(map[byte]bool)
	for i := range b {
		if _, ok := m[b[i]]; ok {
			return true
		}
		m[b[i]] = true
	}
	return false
}