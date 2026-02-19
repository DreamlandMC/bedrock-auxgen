package main

import (
	"github.com/DreamlandMC/bedrock-auxgen/internal/generator"

	"flag"
	"fmt"
	"os"
)

var version = "1.0.1"

func main() {
	out := flag.String("out", "typeIdToAux.json", "output file")
	stdout := flag.Bool("stdout", false, "print only to stdout")
	showVersion := flag.Bool("version", false, "print version")

	flag.Parse()

	if *showVersion {
		fmt.Println("bedrock-auxgen", version)
		return
	}

	auxMap, err := generator.Generate()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if *stdout {
		for k, v := range auxMap {
			fmt.Printf("%s: %d\n", k, v)
		}
		return
	}

	if err := generator.WriteJSON(*out, auxMap); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("✔ generated", *out)
}
