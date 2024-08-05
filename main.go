package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	debugFlag := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	if f, err := os.Open("sqlcube.yaml"); err == nil {
		src, target, err := parseSqlCube(f)
		if err != nil {
			PrintAndExit(err)
		}
		err = CreateTypeAlias(src, target, *debugFlag)
		if err != nil {
			PrintAndExit(err)
		}
		os.Exit(0)
	}

	if f, err := os.Open("sqlc.yaml"); err == nil {
		source, err := parseSqlcYaml(f)
		if err != nil {
			PrintAndExit(err)
		}
		err = CreateTypeAlias(source, source, *debugFlag)
		if err != nil {
			PrintAndExit(err)
		}
		os.Exit(0)
	}

	fmt.Println("Error Neither sqlcube.yaml, nor sqlc.yaml found")
	os.Exit(1)
}

func PrintAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
