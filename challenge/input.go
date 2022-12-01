package challenge

import (
	"bufio"
	"os"
)

type Input struct {
	Lines []string
}

func FromFile(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(f)
}

// func newInputFromReader(f io.Reader, io.Closer) *Input {
// 	result := &Input{

// 	}
// }
