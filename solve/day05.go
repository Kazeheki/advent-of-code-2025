package solve

import "errors"

type Day05Solver struct {
	rawInput []byte
}

func (ds *Day05Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day05Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day05Solver) GetDay() int { return 5 }
