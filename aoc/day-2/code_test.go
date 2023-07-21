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
	you := Player{}
	opp := Player{}
	t.Run("win case", func(t *testing.T) {
		// TODO: Fix this test, I don't think the CalculateScore
		// function is altering the same Player{} that I'm giving it
		yourChoice := "X"
		oppChoice := "C"

		CalculateScore(you, opp, yourChoice, oppChoice)

		want := []Player{{score: 7}, {score: 3}}
		got := []Player{you, opp}

		if !assertSamePlayerSlice(want, got) {
			t.Errorf("got %d wanted %d", got, want)
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

func assertSameStringSlice(sOne, sTwo []string) bool {
	for i, str := range sOne {
		if str != sTwo[i] {
			return false
		}
	}
	return true
}

func assertSamePlayerSlice(sliceOne, sliceTwo []Player) bool {
	for i, player := range sliceOne {
		if player.score != sliceTwo[i].score {
			return false
		}
	}
	return true
}
