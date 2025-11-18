package solve

type CanSolve interface {
	Solve(puzzlePart PuzzlePart) ([]byte, error)
	SetInput(input []byte)
	GetDay() int
}

type PuzzlePart int

const (
	PART_1 PuzzlePart = iota + 1
	PART_2
)
