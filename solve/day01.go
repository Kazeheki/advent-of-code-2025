package solve

import "errors"

type Day01Solver struct {
	rawInput []byte
}

func (ds *Day01Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day01Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day01Solver) GetDay() int { return 1 }
