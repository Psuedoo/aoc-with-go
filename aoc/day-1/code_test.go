package aoc

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("part one", func(t *testing.T) {
		answer := CalculateAnswer("test_data.txt", 1)

		got := answer
		want := 24000

		if got != want {
			t.Errorf("Expected '%d' but got '%d'", want, got)
		}
	})
	t.Run("part two", func(t *testing.T) {
		answer := CalculateAnswer("test_data.txt", 4)

		got := answer
		want := 45000

		if got != want {
			t.Errorf("Expected '%d' but got '%d'", want, got)
		}
	})
}

func TestSortElves(t *testing.T) {
	unsortedElves := Group{
		elves: []Elf{{calorieCount: 3}, {calorieCount: 1}, {calorieCount: 2}},
	}

	sortedElves := Group{
		elves: []Elf{{calorieCount: 1}, {calorieCount: 2}, {calorieCount: 3}},
	}

	unsortedElves.Sort()

	if !assertElvesListSame(unsortedElves, sortedElves) {
		t.Errorf("Expected '%d' but got '%d'", sortedElves, unsortedElves)

	}

}

func TestParseElves(t *testing.T) {
	want := generateTestElves()
	got := parseElves("test_data.txt")

	if !assertElvesListSame(got, want) {
		t.Errorf("Expected '%d' but got '%d'", want, got)
	}
}

func TestCalorieAdd(t *testing.T) {
	elf := Elf{}
	calories := 100

	elf.AddCalories(calories)

	want := 100
	got := elf.calorieCount

	if got != want {
		t.Errorf("Expected '%d' but got '%d'", want, got)
	}
}

func generateTestElves() Group {
	elfList := []Elf{{calorieCount: 6000}, {calorieCount: 4000}, {calorieCount: 11000}, {calorieCount: 24000}, {calorieCount: 10000}}
	elves := Group{}

	for _, elf := range elfList {
		elves.AddElf(elf)
	}

	return elves
}

func assertElvesListSame(elvesOne, elvesTwo Group) bool {
	for index, elf := range elvesOne.elves {
		if elf != elvesTwo.elves[index] {
			return false
		}
	}

	return true
}
