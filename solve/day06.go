package solve

import "errors"

type Day06Solver struct {
	rawInput []byte
}

func (ds *Day06Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day06Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day06Solver) GetDay() int { return 6 }
