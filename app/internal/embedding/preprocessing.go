package embedding

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/bbalet/stopwords"
	"github.com/dchest/stemmer/porter2"
	stripmd "github.com/writeas/go-strip-markdown"
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
	nlTagsRegex := regexp.MustCompile(`['"](?:description|name|title|summary)['"]:\s"([^"]+)"|'([^']+)'`)

	for _, fragment := range fragments {
		var nlFragmentTags []string
		// Find all regex matches
		nlTags := nlTagsRegex.FindAllStringSubmatch(fragment, -1)

		for _, nlTag := range nlTags {
			// Save tag 2 if tag 1 is empty (there are two capturing groups in the regex)
			if strings.Compare(nlTag[1], "") == 0 {
				nlFragmentTags = append(nlFragmentTags, nlTag[2])
			} else {
				nlFragmentTags = append(nlFragmentTags, nlTag[1])
			}
		}

		// Join all extracted strings, remove all Markdown formatting, and append to fragments array
		nlFragments = append(nlFragments, stripmd.Strip(strings.Join(nlFragmentTags, " ")))
	}

	return nlFragments
}

func StopWordRemoval(fragments []string) []string {
	var newFragments []string

	for _, fragment := range fragments {
		newFragments = append(
			newFragments,
			strings.Trim(stopwords.CleanString(fragment, "en", true), " \r\n"),
		)
	}

	return newFragments
}

func Stemming(fragments []string) []string {
	var stemmed []string
	f := func (r rune) bool { return unicode.IsSpace(r) }

	for _, fragment := range fragments {
		var stemmedWords []string
		// Split string into array of words
		words := strings.FieldsFunc(fragment, f)

		for _, word := range words {
			// Perform stemming and append to stemmed words array
			engStemmer := porter2.Stemmer
			//Trim non-alphanumeric characters from strin
			trimmedWord := strings.TrimFunc(word, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})
			stemmedWords = append(stemmedWords, engStemmer.Stem(trimmedWord))
		}

		// Join all words into string and append to stemmed fragments array
		stemmed = append(stemmed, strings.Join(stemmedWords, " "))
	}

	return stemmed
}
