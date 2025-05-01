package gospell

import (
	"errors"
	"io"

	"github.com/zchrykng/gospell/dictionary"
)

type GoSpell struct {
	dictionaries map[string]*dictionary.Dictionary
	defaultDict  string
}

func NewGoSpell() *GoSpell {
	gs := &GoSpell{
		dictionaries: make(map[string]*dictionary.Dictionary),
		defaultDict:  "",
	}

	return gs
}

func (g *GoSpell) AddDictionary(lang, aff, dic string) error {
	if _, prs := g.dictionaries[lang]; prs {
		return errors.New("language already has a dictionary")
	}

	dict, err := dictionary.NewDictionary(aff, dic)
	if err != nil {
		return err
	}

	g.dictionaries[lang] = dict

	if len(g.dictionaries) == 1 {
		g.defaultDict = lang
	}

	return nil
}

func (g *GoSpell) InputConversion(raw []byte) string {
	return g.dictionaries[g.defaultDict].InputConversion(raw)
}

func (g *GoSpell) InputConversionLang(lang string, raw []byte) (string, error) {
	if _, prs := g.dictionaries[lang]; !prs {
		return "", errors.New("selected language dictionary not loaded")
	}
	return g.dictionaries[lang].InputConversion(raw), nil
}

func (g *GoSpell) AddWordRaw(word string) bool {
	return g.dictionaries[g.defaultDict].AddWordRaw(word)
}

func (g *GoSpell) AddWordRawLang(lang string, word string) (bool, error) {
	if _, prs := g.dictionaries[lang]; !prs {
		return false, errors.New("selected language dictionary not loaded")
	}
	return g.dictionaries[lang].AddWordRaw(word), nil
}

func (g *GoSpell) DelWordRaw(word string) bool {
	return g.dictionaries[g.defaultDict].DelWordRaw(word)
}

func (g *GoSpell) DelWordRawLang(lang string, word string) (bool, error) {
	if _, prs := g.dictionaries[lang]; !prs {
		return false, errors.New("selected language dictionary not loaded")
	}
	return g.dictionaries[lang].DelWordRaw(word), nil
}

func (g *GoSpell) AddWordListFile(name string) ([]string, error) {
	return g.dictionaries[g.defaultDict].AddWordListFile(name)
}

func (g *GoSpell) AddWordListFileLang(lang string, name string) ([]string, error) {
	if _, prs := g.dictionaries[lang]; !prs {
		return nil, errors.New("selected language dictionary not loaded")
	}
	return g.dictionaries[lang].AddWordListFile(name)
}

func (g *GoSpell) AddWordList(r io.Reader) ([]string, error) {
	return g.dictionaries[g.defaultDict].AddWordList(r)
}

func (g *GoSpell) AddWordListLang(lang string, r io.Reader) ([]string, error) {
	if _, prs := g.dictionaries[lang]; !prs {
		return nil, errors.New("selected language dictionary not loaded")
	}
	return g.dictionaries[lang].AddWordList(r)
}

func (g *GoSpell) Check(word string) bool {
	return g.dictionaries[g.defaultDict].Check(word)
}

func (g *GoSpell) CheckLang(lang string, word string) (bool, error) {
	if _, prs := g.dictionaries[lang]; !prs {
		return false, errors.New("selected language dictionary not loaded")
	}
	return g.dictionaries[lang].Check(word), nil
}

// Suggest provides suggestions for correcting the spelling of the provided word
func (g *GoSpell) Suggest(word string) []dictionary.Suggestion {
	return g.dictionaries[g.defaultDict].Suggest(word)
}

func (g *GoSpell) SuggestLang(lang string, word string) ([]dictionary.Suggestion, error) {
	if _, prs := g.dictionaries[lang]; !prs {
		return nil, errors.New("selected language dictionary not loaded")
	}
	return g.dictionaries[lang].Suggest(word), nil
}
