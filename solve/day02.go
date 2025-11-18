package solve

import "errors"

type Day02Solver struct {
	rawInput []byte
}

func (ds *Day02Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day02Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day02Solver) GetDay() int { return 2 }
