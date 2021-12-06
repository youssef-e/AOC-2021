package dp2

import (
	"app/pkg/inputReader"
	"strconv"
)

type slidingDepthCalcBoard struct {
	PreviousValue window
	Counter       int
}

type window struct {
	A   int
	B   int
	C   int
	Sum int
}

func DepthSpeed2() (int, error) {

	lines, err := inputReader.ReadInput("DayOne/input.txt")
	if err != nil {
		return 0, err
	}

	calc := slidingDepthCalcBoard{PreviousValue: window{Sum: -1}}
	for i, line := range lines {
		switch i {
		case 0:
			calc.PreviousValue.A, _ = strconv.Atoi(line)
		case 1:
			calc.PreviousValue.B, _ = strconv.Atoi(line)
		case 2:
			calc.PreviousValue.C, _ = strconv.Atoi(line)
			calc.PreviousValue.Sum = calc.PreviousValue.A + calc.PreviousValue.B + calc.PreviousValue.C
		default:
			calc.PreviousValue.A = calc.PreviousValue.B
			calc.PreviousValue.B = calc.PreviousValue.C
			calc.PreviousValue.C, _ = strconv.Atoi(line)
			previousSum := calc.PreviousValue.Sum
			calc.PreviousValue.Sum = calc.PreviousValue.A + calc.PreviousValue.B + calc.PreviousValue.C
			if calc.PreviousValue.Sum > previousSum {
				calc.Counter++
			}
		}

	}
	return calc.Counter, nil
}
