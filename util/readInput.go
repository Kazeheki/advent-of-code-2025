package util

import (
	"fmt"
	"os"
)

const FILE_PATH_FORMAT = "./inputs/day%02d"

func ReadFullDayInput(day int) ([]byte, error) {
	path := fmt.Sprintf(FILE_PATH_FORMAT, day)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}
