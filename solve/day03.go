package solve

import "errors"

type Day03Solver struct {
	rawInput []byte
}

func (ds *Day03Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day03Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day03Solver) GetDay() int { return 3 }
