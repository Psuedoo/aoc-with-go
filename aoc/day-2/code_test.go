package day2

import "testing"

func TestAddScore(t *testing.T) {
	you := Player{}
	opponent := Player{}

	you.AddScore(1)
	opponent.AddScore(3)

	if you.score != 1 || opponent.score != 3 {
		t.Errorf("Expected your score to be %d but got %d\nExpected oppononet score to be %d but got %d", 1, you.score, 3, opponent.score)
	}
}

func TestCalculateScore(t *testing.T) {
	t.Run("win case", func(t *testing.T) {
		want := []*Player{{score: 7}, {score: 3}}
		got := setupTest("X", "C")

		if !assertSamePlayerSlice(want, got) {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
	t.Run("lose case", func(t *testing.T) {
		want := []*Player{{score: 1}, {score: 8}}
		got := setupTest("X", "B")

		if !assertSamePlayerSlice(want, got) {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
	t.Run("tie case", func(t *testing.T) {
		want := []*Player{{score: 4}, {score: 4}}
		got := setupTest("X", "A")

		if !assertSamePlayerSlice(want, got) {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}

func TestRound(t *testing.T) {
	data := readFile("test_data.txt")

	want := []string{"A Y", "B X", "C Z"}

	if !assertSameStringSlice(data, want) {
		t.Errorf("got %v wanted %v", data, want)
	}
}

func setupTest(yourChoice, oppChoice string) []*Player {
	you := &Player{}
	opp := &Player{}

	CalculateScore(you, opp, yourChoice, oppChoice)

	return []*Player{you, opp}
}

func assertSameStringSlice(sOne, sTwo []string) bool {
	for i, str := range sOne {
		if str != sTwo[i] {
			return false
		}
	}
	return true
}

func assertSamePlayerSlice(sliceOne, sliceTwo []*Player) bool {
	for i, player := range sliceOne {
		if player.score != sliceTwo[i].score {
			return false
		}
	}
	return true
}
