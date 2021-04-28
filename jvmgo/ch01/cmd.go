package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	versionFlag bool
	helpFlag    bool
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print usage")
	flag.BoolVar(&cmd.helpFlag, "h", false, "print usage")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version")
	flag.BoolVar(&cmd.versionFlag, "v", false, "print version")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-option] [args] \n", os.Args[0])
}
