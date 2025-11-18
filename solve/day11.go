package solve

import "errors"

type Day11Solver struct {
	rawInput []byte
}

func (ds *Day11Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day11Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day11Solver) GetDay() int { return 11 }
