package question

import "fmt"

const (
	WIDTH  = 600
	HEIGHT = 600
)

// returns question and answer
//
// question should be a 600x600 px image
//
// answers are integers from 1-9 representing correct tiles
type Question func() ([]byte, string, []int, error)

var questions = []Question{pi, color, angle, circle, prime}

func NewQuestion() ([]byte, string, []int, error) {
	n := randint(0, len(questions))
	q := questions[n]
	img, s, ans, err := q()
	if err != nil {
		return nil, "", []int{}, fmt.Errorf("question %d: %w", n, err)
	}
	return img, s, ans, nil
}
