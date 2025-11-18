package util

import (
	"fmt"
	"os"

	"github.com/Kazeheki/advent-of-code-2025/solve"
)

const PATH_FORMAT = "./outputs/day%02d-%d"

func WriteResult(day int, part solve.PuzzlePart, result []byte) error {
	path := fmt.Sprintf(PATH_FORMAT, day, part)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(result)
	if err != nil {
		return err
	}

	return nil
}
