package dictionary

import "fmt"

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

func (s *Suggestion) String() string {
	return fmt.Sprintf("Suggestion[%s](%s)", s.Kind, s.Text)
}

func (d *Dictionary) Suggest(text string) []Suggestion {
	var suggestions []Suggestion

	return suggestions
}
