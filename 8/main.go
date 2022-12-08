package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data := readData("input")
	maxRow := getMaxInRow(data)
	maxCol := getMaxInColumn(data)
	fmt.Println(maxRow)
	fmt.Println(maxCol)
	fmt.Println(getVisibleTrees(data, maxRow, maxCol))

}

func readData(fname string) [][]string{
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	s := bufio.NewScanner(f)
	res := make([][]string, 0)
	for s.Scan() {
		t := s.Text()
		arr := make([]string, 0)
		for i := range t {
			arr = append(arr, string(t[i]))
		}
		res = append(res, arr)
	}
	return res
}
func getVisibleTrees(data [][]string, maxRow, maxCol []string) int {
	c := 0
	for i := 1; i < len(data) - 1; i++ {
		for j := 1; j < len(data[0]) - 1; j++ {
			if data[i][j] < maxCol[i] && data[i][j] < maxRow[j] {
				fmt.Println(data[i][j],i,j)
				c++
			}
		}
	}
	return c
}

func getMaxInRow(in [][]string) []string{
	maxInrows := make([]string, 0)
	for i := range in {
		max := in[i][0]
		for j := range in[i] {
			if in[i][j] > max {
				max = in[i][j]
			}
		}
		maxInrows = append(maxInrows, max)
	}
	return maxInrows
}

func getMaxInColumn(in [][]string) []string{
	maxInCol := make([]string, 0)
	for i := range in[0] {
		max := in[0][i]
		for j := range in {
			if in[j][i] > max {
				max = in[j][i]
			}
		}
		maxInCol = append(maxInCol, max)
	}
	
	return maxInCol
}