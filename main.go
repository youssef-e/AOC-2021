package main

import (
	"app/DayOne/dp1"
	"app/DayOne/dp2"
	"app/DayTwo/p1"
	"app/DayTwo/p2"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

func Run() error {
	var project string
	flag.StringVar(&project, "project", "depthSpeed1", "project : project to launch")
	flag.Parse()

	switch project {
	case "depthSpeed1":
		result, err := dp1.DepthSpeed1()
		if err != nil {
			return err
		}
		fmt.Println(result)
	case "depthSpeed2":
		result, err := dp2.DepthSpeed2()
		if err != nil {
			return err
		}
		fmt.Println(result)
	case "area1":
		result, err := p1.Area()
		if err != nil {
			return err
		}
		fmt.Println(result)
	case "area2":
		result, err := p2.Area()
		if err != nil {
			return err
		}
		fmt.Println(result)
	default:
		return fmt.Errorf("invalid command")
	}
	return nil
}
