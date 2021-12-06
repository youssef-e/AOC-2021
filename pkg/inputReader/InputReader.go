package inputReader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(path string) ([]string, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to open %s", path)

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	return text, nil
}
