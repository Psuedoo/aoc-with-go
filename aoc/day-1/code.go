package aoc

import (
	"bufio"
	"fmt"
	"os"
)

type Elf struct {
	calorieCount int
}

func (e *Elf) AddCalories(calories int) {
	e.calorieCount += calories
}

type Group struct {
	elves []Elf
}

func (g *Group) AddElf(elf Elf) {
	g.elves = append(g.elves, elf)
}

func (g *Group) Sort() {
	g.elves = g.elves
}

func CalculateAnswer(filename string) int {
	return 10
}

func parseElves(filename string) Group {

	var calories int

	data := readFile(filename)
	firstElf := Elf{}
	elves := Group{}

	currentElf := firstElf

	for _, cal := range data {
		calories = 0
		if cal == "" {
			elves.AddElf(currentElf)
			currentElf = Elf{}
		} else {
			fmt.Sscan(cal, &calories)
			currentElf.AddCalories(calories)
		}
	}

	return elves
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

func main() {
	print("Hello")
}
