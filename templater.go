package main

import (
	"templater"
	"flag"
)

func main() {
	var filepath = flag.String("filepath", ".", "path to write the generated json")
	flag.Parse()
	templater.GenerateTemplates(*filepath)
}
