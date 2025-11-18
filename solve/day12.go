package solve

import "errors"

type Day12Solver struct {
	rawInput []byte
}

func (ds *Day12Solver) SetInput(input []byte) {
	ds.rawInput = input
}

func (ds *Day12Solver) Solve(puzzlePart PuzzlePart) ([]byte, error) {
	return nil, errors.New("not yet implemented")
}

func (ds Day12Solver) GetDay() int { return 12 }
