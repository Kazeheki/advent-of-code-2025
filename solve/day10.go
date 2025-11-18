package solve

import "errors"

type Day10Solver struct {
	rawInput []byte
}

func (ds *Day10Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day10Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day10Solver) GetDay() int { return 10 }
