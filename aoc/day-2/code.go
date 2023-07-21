package day2

import (
	"bufio"
	"fmt"
	"os"
)

var pointMap = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
	"lost":     0,
	"tied":     3,
	"won":      6,
}

var choiceMap = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var oppositeMap = map[string]string{
	"tied": "tied",
	"won":  "lost",
	"lost": "won",
}

var winMap = map[string]string{
	"AX": "tied",
	"AY": "won",
	"AZ": "lost",
	"BX": "lost",
	"BY": "tied",
	"BZ": "won",
	"CX": "won",
	"CY": "lost",
	"CZ": "tied",
}

type Player struct {
	score int
}

func (p *Player) AddScore(number int) {
	p.score += number
}

func CalculateScore(playerOne, playerTwo Player, pOneChoice, pTwoChoice string) {
	outcome := winMap[pTwoChoice+pOneChoice]
	pOneNormChoice := choiceMap[pOneChoice]
	pTwoNormChoice := choiceMap[pTwoChoice]

	playerOne.AddScore(pointMap[outcome])
	playerOne.AddScore(pointMap[pOneNormChoice])

	playerTwo.AddScore(pointMap[oppositeMap[outcome]])
	playerTwo.AddScore(pointMap[pTwoNormChoice])
}

func readFile(filename string) []string {
	var data []string
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		data = append(data, fileScanner.Text())
	}

	file.Close()

	return data
}
