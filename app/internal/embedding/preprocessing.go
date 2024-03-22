package embedding

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/bbalet/stopwords"
	"github.com/dchest/stemmer/porter2"
	stripmd "github.com/writeas/go-strip-markdown"
)

// PreprocessFragment - fragments are preprocessed by running a standard NLP pipeline, composed of string cleaning,
// stop-word removal, and stemming. An array of fragments (string), and a boolean indicating is the fragments are
// queries or not need to be passed to the function.
func PreprocessFragment(fragments []string, isQuery bool) []string {
	cleanFragment := fragments

	if !isQuery {
		cleanFragment = ExtractTags(cleanFragment)
	}

	return Stemming(StopWordRemoval(cleanFragment))
}

// ExtractTags - extract the NL tags from a fragment (JSON document documenting a REST API), and return an array of
// strings, one for each fragment. An array of fragments needs to be passed to the function.
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

// StopWordRemoval - remove all stopwords from the given strings. An array of strings needs to be passed to the
// function.
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

// Stemming - stem all the words contained in the given strings. An array of strings needs to be passed to the
// function.
func Stemming(fragments []string) []string {
	var stemmed []string
	f := func(r rune) bool { return unicode.IsSpace(r) }

	for _, fragment := range fragments {
		var stemmedWords []string
		// Split string into array of words
		words := strings.FieldsFunc(fragment, f)

		for _, word := range words {
			// Perform stemming and append to stemmed words array
			engStemmer := porter2.Stemmer
			//Trim non-alphanumeric characters from string
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
