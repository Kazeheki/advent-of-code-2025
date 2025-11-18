package solve

import "errors"

type Day08Solver struct {
	rawInput []byte
}

func (ds *Day08Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day08Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day08Solver) GetDay() int { return 8 }
