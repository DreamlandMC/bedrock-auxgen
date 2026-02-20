package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DreamlandMC/bedrock-auxgen/internal/generator"
)

var version = "1.0.1"

func main() {
	out := flag.String("out", "typeIdToAux.json", "output file")
	stdout := flag.Bool("stdout", false, "print only to stdout")
	showVersion := flag.Bool("version", false, "print version")
	customRP := flag.String("rp", "", "path to resource pack root")
	customStart := flag.Int64("custom-start", 257, "starting ID for custom items (default: 257)")

	flag.Parse()

	if *showVersion {
		fmt.Println("bedrock-auxgen", version)
		return
	}

	auxMap, err := generator.Generate(*customRP, *customStart)
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