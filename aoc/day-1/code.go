package aoc

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	sort.Slice(g.elves, func(i, j int) bool {
		return g.elves[i].calorieCount > g.elves[j].calorieCount
	})
}

func CalculateAnswer(filename string, topCount int) int {
	group := parseElves(filename)
	group.Sort()
	answer := 0
	for i := 0; i < topCount; i++ {
		answer += group.elves[i].calorieCount
	}

	return answer
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
	answer := CalculateAnswer("data.txt", 1)
	fmt.Printf("Answer: %d", answer)
}
