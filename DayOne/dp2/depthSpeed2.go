package dp2

import (
	"app/pkg/inputReader"
	"fmt"
	"strconv"
)

type slidingDepthCalcBoard struct {
	PreviousValue string
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

	calc := slidingDepthCalcBoard{}
	for _, line := range lines {

		if calc.PreviousValue != "" {
			current, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			previous, err := strconv.Atoi(calc.PreviousValue)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			if current > previous {
				calc.Counter++
			}
		}

		calc.PreviousValue = line
	}
	return calc.Counter, nil
}
