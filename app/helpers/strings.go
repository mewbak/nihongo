package helpers

import (
	"github.com/gojp/kana"
)

// wrap a string in strong tags
func MakeStrong(query string) string {
	return "<strong>" + query + "</strong>"
}

// convert the query to hiragana and katakana. if it's already in
// hiragana or katakana, it will just be the same.
func ConvertQueryToKana(query string) (hiragana, katakana string) {
	h := kana.RomajiToHiragana(query)
	k := kana.RomajiToKatakana(query)
	return h, k
}