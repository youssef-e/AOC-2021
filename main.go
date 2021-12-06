package main

import (
	"app/DayOne/dp1"
	"app/DayOne/dp2"
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
	default:
		return fmt.Errorf("invalid command")
	}
	return nil
}
