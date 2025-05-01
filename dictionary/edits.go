package dictionary

func (d *Dictionary) edit1(word string) []string {
	var result []string

	wordChars := d.WordChars()

	result = append(result, split(word)...)

	return nil
}

func split(word string) []string {
	var results []string

	for i := 1; i < len(word); i++ {
		results = append(results, word[:i], word[i:])
	}

	return results
}

func deletes(word string) []string {
	var results []string

	for i := 0; i < len(word); i++ {
		results = append(results, word[:i]+word[i+1:])
	}

	return results
}

//def edits1(word):
//    "All edits that are one edit away from `word`."
//    letters    = 'abcdefghijklmnopqrstuvwxyz'
//    deletes    = [L + R[1:]               for L, R in splits if R]
//    transposes = [L + R[1] + R[0] + R[2:] for L, R in splits if len(R)>1]
//    replaces   = [L + c + R[1:]           for L, R in splits if R for c in letters]
//    inserts    = [L + c + R               for L, R in splits for c in letters]
//    return set(deletes + transposes + replaces + inserts)
//
//def edits2(word):
//    "All edits that are two edits away from `word`."
//    return (e2 for e1 in edits1(word) for e2 in edits1(e1))
