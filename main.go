package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {

	debugFlag := flag.Bool("debug", false, "enable debug mode")
	if len(os.Args) < 2 {
		PrintAndExit(errors.New("provide argument generate or reduce"))
	}

	command := os.Args[1]

	flag.Parse()
	f, err := os.Open("sqlcube.yaml")
	if err != nil {
		PrintAndExit(fmt.Errorf("sqlcube.yaml open error: %s", err.Error()))
	}
	cfg, err := parseSqlCube(f)
	PrintAndExit(err)

	switch command {
	case "reduce":
		err = CreateTypeAlias(cfg.Go.Source, cfg.Go.Target, *debugFlag)
		PrintAndExit(err)
	case "generate":
		err = GenerateSqlc(cfg.Generation)
		PrintAndExit(err)
	default:
		PrintAndExit(fmt.Errorf("unknown command %s", command))
	}

	os.Exit(0)
}

func PrintAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
