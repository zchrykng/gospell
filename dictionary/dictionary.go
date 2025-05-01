package dictionary

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type HashMap map[string]int

// Dictionary is main struct
type Dictionary struct {
	Config DictConfig
	Dict   map[string]struct{} // likely will contain some value later
	Hash   map[string]HashMap

	ireplacer *strings.Replacer // input conversion
	compounds []*regexp.Regexp
}

// InputConversion does any character substitution before checking
//
//	This is based on the ICONV stanza
func (d *Dictionary) InputConversion(raw []byte) string {
	sraw := string(raw)
	if d.ireplacer == nil {
		return sraw
	}
	return d.ireplacer.Replace(sraw)
}

// AddWordRaw adds a single word to the internal dictionary without modifications
// returns true if added
// return false is already exists
func (d *Dictionary) AddWordRaw(word string) bool {
	_, ok := d.Dict[word]
	if ok {
		// already exists
		return false
	}
	d.Dict[word] = struct{}{}
	return true
}

func (d *Dictionary) DelWordRaw(word string) bool {
	_, ok := d.Dict[word]
	if ok {
		delete(d.Dict, word)
		return true
	}

	return false
}

// AddWordListFile reads in a word list file
func (d *Dictionary) AddWordListFile(name string) ([]string, error) {
	fd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	return d.AddWordList(fd)
}

// AddWordList adds basic word lists, just one word per line
//
//	Assumed to be in UTF-8
//
// TODO: hunspell compatible with "*" prefix for forbidden words
// and affix support
// returns list of duplicated words and/or error
func (d *Dictionary) AddWordList(r io.Reader) ([]string, error) {
	var duplicates []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line == "#" {
			continue
		}
		for _, word := range CaseVariations(line, CaseStyle(line)) {
			if !d.AddWordRaw(word) {
				duplicates = append(duplicates, word)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return duplicates, err
	}
	return duplicates, nil
}

// Check to see if a given word is in the internal dictionaries
// TODO: add multiple dictionaries
func (d *Dictionary) Check(word string) bool {
	// log.Printf("Checking %s", word)
	_, ok := d.Dict[word]
	if ok {
		return true
	}
	if isNumber(word) {
		return true
	}
	if isNumberHex(word) {
		return true
	}

	if isNumberBinary(word) {
		return true
	}

	if isHash(word) {
		return true
	}

	// check compounds
	for _, pat := range d.compounds {
		if pat.MatchString(word) {
			return true
		}
	}

	// Maybe a word with units? e.g. 100GB
	units := isNumberUnits(word)
	if units != "" {
		// dictionary appears to have list of units
		if _, ok = d.Dict[units]; ok {
			return true
		}
	}

	// if camelCase and each word e.g. "camel" "Case" is know
	// then the word is considered known
	if chunks := splitCamelCase(word); len(chunks) > 0 {
		if false {
			for _, chunk := range chunks {
				if _, ok = d.Dict[chunk]; !ok {
					return false
				}
			}
		}
		return true
	}

	return false
}

func (d *Dictionary) WordChars() string {
	return d.Config.WordChars
}

// NewDictionaryReader creates a speller from io.Readers for
// Hunspell files
func NewDictionaryReader(aff, dic io.Reader) (*Dictionary, error) {
	affix, err := NewDictConfig(aff)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(dic)
	// get first line
	if !scanner.Scan() {
		return nil, scanner.Err()
	}
	line := scanner.Text()
	i, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		return nil, err
	}

	gs := Dictionary{
		Dict:      make(map[string]struct{}, i*5),
		Hash:      make(map[string]HashMap, i*5),
		compounds: make([]*regexp.Regexp, 0, len(affix.CompoundRule)),
	}

	var words []string
	for scanner.Scan() {
		line := scanner.Text()
		words, err = affix.Expand(line, words)
		if err != nil {
			return nil, fmt.Errorf("unable to process %q: %s", line, err)
		}

		if len(words) == 0 {
			// log.Printf("No words for %s", line)
			continue
		}

		style := CaseStyle(words[0])
		for _, word := range words {
			for _, wordform := range CaseVariations(word, style) {
				gs.Dict[wordform] = struct{}{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for _, compoundRule := range affix.CompoundRule {
		pattern := "^"
		for _, key := range compoundRule {
			switch key {
			case '(', ')', '+', '?', '*':
				pattern += string(key)
			default:
				groups := affix.compoundMap[key]
				pattern = pattern + "(" + strings.Join(groups, "|") + ")"
			}
		}
		pattern += "$"
		pat, err := regexp.Compile(pattern)
		if err != nil {
			log.Printf("REGEXP FAIL= %q %s", pattern, err)
		} else {
			gs.compounds = append(gs.compounds, pat)
		}

	}

	if len(affix.IconvReplacements) > 0 {
		gs.ireplacer = strings.NewReplacer(affix.IconvReplacements...)
	}
	return &gs, nil
}

// NewDictionary from AFF and DIC Hunspell filenames
func NewDictionary(affFile, dicFile string) (*Dictionary, error) {
	aff, err := os.Open(affFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open aff: %s", err)
	}
	defer aff.Close()
	dic, err := os.Open(dicFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open dic: %s", err)
	}
	defer dic.Close()
	h, err := NewDictionaryReader(aff, dic)
	return h, err
}
