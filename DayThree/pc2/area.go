package p2

import (
	"app/pkg/inputReader"
	"strconv"
	"strings"
)

type coord struct {
	X   int
	Y   int
	Aim int
}

func Area() (int, error) {

	lines, err := inputReader.ReadInput("DayTwo/input.txt")
	if err != nil {
		return 0, err
	}
	pos := coord{}
	for _, line := range lines {
		command := strings.Split(line, " ")
		value, err := strconv.Atoi(command[1])
		if err != nil {
			return 0, err
		}
		switch command[0] {
		case "forward":
			pos.X += value
			pos.Y += value * pos.Aim
		case "up":
			pos.Aim -= value
		case "down":
			pos.Aim += value
		}
		if pos.Y < 0 {
			pos.Y = 0
		}
	}
	return (pos.X * pos.Y), nil
}
