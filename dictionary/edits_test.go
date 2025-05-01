package dictionary

import (
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

func TestSplit(t *testing.T) {
	word := "test"
	expected := []pair{
		{"", "test"},
		{"t", "est"},
		{"te", "st"},
		{"tes", "t"},
		{"test", ""},
	}

	result := split(word)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("split(%q) = %v; want %v", word, result, expected)
	}
}

func TestDeletes(t *testing.T) {
	pairs := split("test")

	expected := mapset.NewSet[string]()
	expected.Append("est", "tst", "tet", "tes")

	result := mapset.NewSet[string]()
	result.Append(deletes(pairs)...)

	if !expected.Equal(result) {
		t.Errorf("deletes(%v) = %v; want %v", pairs, result, expected)
	}
}

func TestTransposes(t *testing.T) {
	pairs := split("test")

	expected := mapset.NewSet[string]()
	expected.Append("tets", "etst", "tset", "tets")

	result := mapset.NewSet[string]()
	result.Append(transposes(pairs)...)

	if !expected.Equal(result) {
		t.Errorf("transposes(%v) = %v; want %v", pairs, result, expected)
	}
}

func TestReplacements(t *testing.T) {
	pairs := split("test")

	characters := "abc"
	expected := mapset.NewSet[string]()
	expected.Append(
		"aest", "best", "cest",
		"tast", "tbst", "tcst",
		"teat", "tebt", "tect",
		"tesa", "tesb", "tesc")

	result := mapset.NewSet[string]()
	result.Append(replacements(pairs, characters)...)

	if !expected.Equal(result) {
		t.Errorf("replacements(%v, %q) = %v; want %v", pairs, characters, result, expected)
	}
}

func TestInserts(t *testing.T) {
	pairs := split("test")

	characters := "abc"

	expected := mapset.NewSet[string]()
	expected.Append(
		"atest", "btest", "ctest",
		"taest", "tbest", "tcest",
		"teast", "tebst", "tecst",
		"tesat", "tesbt", "tesct",
		"testa", "testb", "testc",
	)

	result := mapset.NewSet[string]()

	result.Append(inserts(pairs, characters)...)

	if !expected.Equal(result) {
		t.Errorf("inserts(%v, %q) = %v; want %v", pairs, characters, result, expected)
	}
}

func TestEdit1(t *testing.T) {
	word := "test"
	characters := "abc"

	expected := mapset.NewSet[string]()
	expected.Append("est", "tst", "tet", "tes")
	expected.Append("tets", "etst", "tset", "tets")
	expected.Append(
		"aest", "best", "cest",
		"tast", "tbst", "tcst",
		"teat", "tebt", "tect",
		"tesa", "tesb", "tesc",
	)
	expected.Append(
		"atest", "btest", "ctest",
		"taest", "tbest", "tcest",
		"teast", "tebst", "tecst",
		"tesat", "tesbt", "tesct",
		"testa", "testb", "testc",
	)

	result := mapset.NewSet[string]()
	result.Append(edit1(word, characters)...)

	if !expected.Equal(result) {
		t.Errorf("edit1(%q) returned no results; expected non-empty result", word)
	}
}

func TestEdit2(t *testing.T) {
	word := "test"
	characters := "abc"

	expected := mapset.NewSet[string]()
	expected.Append("est", "tst", "tet", "tes")
	expected.Append("tets", "etst", "tset", "tets")
	expected.Append(
		"aest", "best", "cest",
		"tast", "tbst", "tcst",
		"teat", "tebt", "tect",
		"tesa", "tesb", "tesc",
	)
	expected.Append(
		"atest", "btest", "ctest",
		"taest", "tbest", "tcest",
		"teast", "tebst", "tecst",
		"tesat", "tesbt", "tesct",
		"testa", "testb", "testc",
	)

	result := mapset.NewSet[string]()
	result.Append(edit2(word, characters)...)

	if result.IsEmpty() {
		t.Errorf("edit2(%q) returned no results; expected non-empty result", word)
	} else if !expected.Equal(result) {
		t.Errorf("edit2(%q) = %v; want %v", word, result, expected)
	}

}
