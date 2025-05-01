package dictionary

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type pair struct {
	left  string
	right string
}

func edit1(word string, characters string) []string {
	results := mapset.NewSet[string]()

	splits := split(word)

	results.Append(deletes(splits)...)
	results.Append(transposes(splits)...)
	results.Append(replacements(splits, characters)...)
	results.Append(inserts(splits, characters)...)

	return results.ToSlice()
}

func edit2(word string, characters string) []string {
	//def edits2(word):
	//    "All edits that are two edits away from `word`."
	//    return (e2 for e1 in edits1(word) for e2 in edits1(e1))
	results := mapset.NewSet[string]()

	for _, edit := range edit1(word, characters) {
		results.Append(edit1(edit, characters)...)
	}

	return results.ToSlice()
}

func split(word string) []pair {
	var results []pair

	for i := 0; i <= len(word); i++ {
		results = append(results, pair{word[:i], word[i:]})
	}

	return results
}

func deletes(pairs []pair) []string {
	var results []string

	for _, p := range pairs {
		if len(p.right) > 0 {
			results = append(results, p.left+p.right[1:])
		}
	}

	return results
}

func transposes(pairs []pair) []string {
	var results []string

	for _, p := range pairs {
		if len(p.right) > 1 {
			results = append(results, p.left+string(p.right[1])+string(p.right[0])+p.right[2:])
		}
	}

	return results
}

func replacements(pairs []pair, characters string) []string {
	var results []string

	for _, p := range pairs {
		if len(p.right) > 0 {
			for _, c := range characters {
				results = append(results, p.left+string(c)+p.right[1:])
			}
		}
	}

	return results
}

func inserts(pairs []pair, characters string) []string {
	var results []string

	for _, p := range pairs {
		for _, c := range characters {
			results = append(results, p.left+string(c)+p.right)
		}
	}

	return results
}
