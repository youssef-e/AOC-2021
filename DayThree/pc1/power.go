package pc1

import (
	"app/pkg/inputReader"
	"math"
)

type bit struct {
	Ones  int
	Zeros int
}

func Area() (int, error) {

	lines, err := inputReader.ReadInput("DayTwo/input.txt")
	if err != nil {
		return 0, err
	}
	consumption := map[int]bit{}
	for _, line := range lines {
		for i := 0; i < 5; i++ {
			index := math.Mod(float64(i), 5)
			if string(line[int(index)] == string("0"[0])) {
				consumption[index].Zeros++
			}
		}

	}
	return 0, nil
}
