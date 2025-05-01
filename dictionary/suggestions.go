package dictionary

import (
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
)

func NewSuggestion(kind, text string) Suggestion {
	return Suggestion{
		Kind: kind,
		Text: text,
	}
}

type Suggestion struct {
	Text string
	Kind string
}

type result struct {
	Word  string
	Count int
}

func orderPossibilities(possibilities HashMap) []string {
	var results []result
	for word, count := range possibilities {
		results = append(results, result{Word: word, Count: count})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Count > results[j].Count
	})

	var ordered []string
	for _, r := range results {
		ordered = append(ordered, r.Word)
	}
	return ordered
}

func (s *Suggestion) String() string {
	return fmt.Sprintf("Suggestion[%s](%s)", s.Kind, s.Text)
}

func (d *Dictionary) Suggest(text string) []Suggestion {
	var suggestions []Suggestion

	results := mapset.NewSet[string]()

	if d.Check(text) {
		results.Add(text)
	}

	for _, edit := range edit1(text, d.WordChars()) {
		if d.Check(edit) && !results.Contains(edit) {
			results.Add(edit)
			suggestions = append(suggestions, NewSuggestion("edit1", edit))
		}
	}

	for _, edit2 := range edit2(text, d.WordChars()) {
		if d.Check(edit2) && !results.Contains(edit2) {
			results.Add(edit2)
			suggestions = append(suggestions, NewSuggestion("edit2", edit2))
		}
	}

	return suggestions
}
