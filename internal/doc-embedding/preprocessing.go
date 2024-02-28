package doc_embedding

import (
	stripmd "github.com/writeas/go-strip-markdown"
	"regexp"
	"strings"
	"unicode"

	"github.com/bbalet/stopwords"
	"github.com/dchest/stemmer/porter2"
)


func PreprocessFragment(fragments []string, isQuery bool) []string {
	cleanFragment := fragments

	if !isQuery {
		cleanFragment = ExtractTags(cleanFragment)
	}

	return Stemming(StopWordRemoval(cleanFragment))
}

func ExtractTags(fragments []string) []string {
	var nlFragments []string
	nlTagsRegex := regexp.MustCompile(`['"](?:description|name|title|summary)['"]:\s['"]([^'"]+)['"]`)

	for _, fragment := range fragments {
		var nlFragmentTags []string
		nlTags := nlTagsRegex.FindAllStringSubmatch(fragment, -1)

		for _, nlTag := range nlTags {
			nlFragmentTags = append(nlFragmentTags, nlTag[1])
		}

		nlFragments = append(nlFragments, stripmd.Strip(strings.Join(nlFragmentTags, " ")))
	}

	return nlFragments
}


func StopWordRemoval(fragments []string) []string {
	var newFragments []string

	for _, fragment := range fragments {
		newFragments = append(newFragments, stopwords.CleanString(fragment, "en", true))
	}

	return newFragments
}

func Stemming(fragments []string) []string {
	var stemmed []string
	f := func (char rune) bool { return unicode.IsSpace(char) }

	for _, fragment := range fragments {
		var stemmedWords []string
		words := strings.FieldsFunc(fragment, f)

		for _, word := range words {
			engStemmer := porter2.Stemmer
			stemmedWords = append(stemmedWords, engStemmer.Stem(word))
		}

		stemmed = append(stemmed, strings.Join(stemmedWords, " "))
	}

	return stemmed
}
