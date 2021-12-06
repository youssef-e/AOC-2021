package dp1

import (
	"app/pkg/inputReader"
	"fmt"
	"strconv"
)

type depthCalcBoard struct {
	PreviousValue string
	Counter       int
}

func DepthSpeed1() (int, error) {

	lines, err := inputReader.ReadInput("DayOne/input.txt")
	if err != nil {
		return 0, err
	}

	calc := depthCalcBoard{}
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
