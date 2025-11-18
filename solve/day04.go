package solve

import "errors"

type Day04Solver struct {
	rawInput []byte
}

func (ds *Day04Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day04Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day04Solver) GetDay() int { return 4 }
