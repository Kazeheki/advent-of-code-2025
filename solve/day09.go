package solve

import "errors"

type Day09Solver struct {
	rawInput []byte
}

func (ds *Day09Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day09Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day09Solver) GetDay() int { return 9 }
