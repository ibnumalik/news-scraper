package main

import (
	"fmt"
	"os"

	"github.com/ibnumalik/amanz-scraper/amanz"
)

func usage() {
	fmt.Println("Usage:\n news [newssource] [newstype]\n\n" +
		"By default [newstype] is latest, so you can just give [newssource]\n" +
		"news amanz\n\n" +
		"[newssource]\tcan be any of this: amanz\n" +
		"[newstype]\tcan be any one of this: trending, latest, featured\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		usage()
	}

	source := os.Args[1]

	types := "latest"

	if (len(os.Args)) >= 3 {
		types = os.Args[2]
	}

	switch source {
	case "amanz":
		amanz.Get(types)
	default:
		fmt.Println("The news source is not supported yet.")
		return
	}
}
