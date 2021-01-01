package gofakeit

import (
	"math/rand"
	"strings"
)

// HackerPhrase will return a random hacker sentence
func HackerPhrase() string { return hackerPhrase(globalFaker.Rand) }

// HackerPhrase will return a random hacker sentence
func (f *Faker) HackerPhrase() string { return hackerPhrase(f.Rand) }

func hackerPhrase(r *rand.Rand) string {
	words := strings.Split(Generate(getRandValue(r, []string{"hacker", "phrase"})), " ")
	words[0] = strings.Title(words[0])
	return strings.Join(words, " ")
}

// HackerAbbreviation will return a random hacker abbreviation
func HackerAbbreviation() string { return hackerAbbreviation(globalFaker.Rand) }

// HackerAbbreviation will return a random hacker abbreviation
func (f *Faker) HackerAbbreviation() string { return hackerAbbreviation(f.Rand) }

func hackerAbbreviation(r *rand.Rand) string {
	return getRandValue(r, []string{"hacker", "abbreviation"})
}

// HackerAdjective will return a random hacker adjective
func HackerAdjective() string { return hackerAdjective(globalFaker.Rand) }

// HackerAdjective will return a random hacker adjective
func (f *Faker) HackerAdjective() string { return hackerAdjective(f.Rand) }

func hackerAdjective(r *rand.Rand) string {
	return getRandValue(globalFaker.Rand, []string{"hacker", "adjective"})
}

// HackerNoun will return a random hacker noun
func HackerNoun() string { return hackerNoun(globalFaker.Rand) }

// HackerNoun will return a random hacker noun
func (f *Faker) HackerNoun() string { return hackerNoun(f.Rand) }

func hackerNoun(r *rand.Rand) string {
	return getRandValue(r, []string{"hacker", "noun"})
}

// HackerVerb will return a random hacker verb
func HackerVerb() string { return hackerVerb(globalFaker.Rand) }

// HackerVerb will return a random hacker verb
func (f *Faker) HackerVerb() string { return hackerVerb(f.Rand) }

func hackerVerb(r *rand.Rand) string {
	return getRandValue(globalFaker.Rand, []string{"hacker", "verb"})
}

// HackeringVerb will return a random hacker ingverb
func HackeringVerb() string { return hackeringVerb(globalFaker.Rand) }

// HackeringVerb will return a random hacker ingverb
func (f *Faker) HackeringVerb() string { return hackeringVerb(f.Rand) }

func hackeringVerb(r *rand.Rand) string {
	return getRandValue(globalFaker.Rand, []string{"hacker", "ingverb"})
}

func addHackerLookup() {
	AddFuncLookup("hackerphrase", Info{
		Display:     "Hacker Phrase",
		Category:    "hacker",
		Description: "Random hacker phrase",
		Example:     "If we calculate the program, we can get to the AI pixel through the redundant XSS matrix!",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerPhrase(), nil
		},
	})

	AddFuncLookup("hackerabbreviation", Info{
		Display:     "Hacker Abbreviation",
		Category:    "hacker",
		Description: "Random hacker abbreviation",
		Example:     "ADP",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerAbbreviation(), nil
		},
	})

	AddFuncLookup("hackeradjective", Info{
		Display:     "Hacker Adjective",
		Category:    "hacker",
		Description: "Random hacker adjective",
		Example:     "wireless",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerAdjective(), nil
		},
	})

	AddFuncLookup("hackernoun", Info{
		Display:     "Hacker Noun",
		Category:    "hacker",
		Description: "Random hacker noun",
		Example:     "driver",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerNoun(), nil
		},
	})

	AddFuncLookup("hackerverb", Info{
		Display:     "Hacker Verb",
		Category:    "hacker",
		Description: "Random hacker verb",
		Example:     "synthesize",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerVerb(), nil
		},
	})

	AddFuncLookup("hackeringverb", Info{
		Display:     "Hackering Verb",
		Category:    "hacker",
		Description: "Random hackering verb",
		Example:     "connecting",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackeringVerb(), nil
		},
	})
}
