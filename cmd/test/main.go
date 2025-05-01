package main

import (
	"flag"
	"fmt"

	"github.com/zchrykng/gospell"
)

func main() {
	aff := flag.String("aff", "cmd/test/en_US.aff", "the aff file for the dictionary to import")
	dic := flag.String("dic", "cmd/test/en_US.dic", "the dic file for the dictionary to import")
	lang := flag.String("lang", "en_US", "the language of the imported dictionary")

	flag.Parse()

	//if os.Stat(flag.aff)

	gs := gospell.NewGoSpell()

	err := gs.AddDictionary(*lang, *aff, *dic)
	if err != nil {
		panic(err)
	}

	words := []string{
		"blah",
		"bopqiwne",
		"hello",
		"brown",
		"green",
		"yellow",
		"silver",
		"basalt",
		"ardvark",
		"aardvark",
		"aardvarks",
		"aardvarkes",
		"aardvarker",
		"aardvarking",
		"run",
		"runner",
		"running",
		"CamelCase",
		"joint venture",
	}

	for _, word := range words {
		fmt.Printf("test: %s, result: %v\n", word, gs.Check(word))
	}

	word := "toast"
	suggestions := gs.Suggest(word)
	fmt.Printf("test: %s, suggestions: %v\n", word, suggestions)
}
