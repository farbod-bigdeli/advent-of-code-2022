package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type RPS struct {
	FirstPlayer  string
	SecondPlayer string
	Score        int
	Score2       int
}

const (
	FPlayerRock     = "A"
	FPlayerPaper    = "B"
	FPlayerScissors = "C"
	SPlayerRock     = "X"
	SPlayerPaper    = "Y"
	SPlayerScissor  = "Z"


)

func main()  {
	rps := readData("input")
	score := 0
	score2 := 0
	for _, v := range(rps) {
		score += v.Score
		score2 += v.Score2
	}
	fmt.Println(score, score2)
}

func readData(fname string) ([]RPS) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open file: %s", err.Error())
	}
	s := bufio.NewScanner(f)
	res := make([]RPS, 0)
	for s.Scan() {
		text := s.Text()
		inpArr := strings.Split(text, " ")
		rps := RPS{FirstPlayer: inpArr[0], SecondPlayer: inpArr[1]}
		rps.calculateScore()
		res = append(res, rps)
	}
	return res
}


func (in *RPS) calculateScore(){
	in.Score = in.logicScore()
	in.Score2 = in.logicScore2()
}



func (in *RPS) logicScore() int {
	if (in.FirstPlayer == FPlayerRock) && (in.SecondPlayer == SPlayerRock) {
		return 3 + 1
	} else if (in.FirstPlayer == FPlayerRock) && (in.SecondPlayer == SPlayerPaper) {
		return 6 + 2
	} else if (in.FirstPlayer == FPlayerRock) && (in.SecondPlayer == SPlayerScissor) {
		return 0 + 3
	} else if (in.FirstPlayer == FPlayerPaper) && (in.SecondPlayer == SPlayerRock) {
		return 0 + 1
	} else if (in.FirstPlayer == FPlayerPaper) && (in.SecondPlayer == SPlayerPaper) {
		return 3 + 2
	} else if (in.FirstPlayer == FPlayerPaper) && (in.SecondPlayer == SPlayerScissor) {
		return 6 + 3
	} else if (in.FirstPlayer == FPlayerScissors) && (in.SecondPlayer == SPlayerRock) {
		return 6 + 1
	} else if (in.FirstPlayer == FPlayerScissors) && (in.SecondPlayer == SPlayerPaper) {
		return 0 + 2
	} else if (in.FirstPlayer == FPlayerScissors) && (in.SecondPlayer == SPlayerScissor) {
		return 3 + 3
	} else {
		log.Fatal("wrong input")
		return 0 
	}
}

func (in *RPS) logicScore2() int {
	if (in.FirstPlayer == FPlayerRock) && (in.SecondPlayer == SPlayerRock) {
		return 0 + 3
	} else if (in.FirstPlayer == FPlayerRock) && (in.SecondPlayer == SPlayerPaper) {
		return 3 + 1
	} else if (in.FirstPlayer == FPlayerRock) && (in.SecondPlayer == SPlayerScissor) {
		return 6 + 2
	} else if (in.FirstPlayer == FPlayerPaper) && (in.SecondPlayer == SPlayerRock) {
		return 0 + 1
	} else if (in.FirstPlayer == FPlayerPaper) && (in.SecondPlayer == SPlayerPaper) {
		return 3 + 2
	} else if (in.FirstPlayer == FPlayerPaper) && (in.SecondPlayer == SPlayerScissor) {
		return 6 + 3
	} else if (in.FirstPlayer == FPlayerScissors) && (in.SecondPlayer == SPlayerRock) {
		return 0 + 2
	} else if (in.FirstPlayer == FPlayerScissors) && (in.SecondPlayer == SPlayerPaper) {
		return 3 + 3
	} else if (in.FirstPlayer == FPlayerScissors) && (in.SecondPlayer == SPlayerScissor) {
		return 6 + 1
	} else {
		log.Fatal("wrong input")
		return 0 
	}
	
}
