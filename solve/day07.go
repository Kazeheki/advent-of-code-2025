package solve

import "errors"

type Day07Solver struct {
	rawInput []byte
}

func (ds *Day07Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day07Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day07Solver) GetDay() int { return 7 }
