package utils

import (
	"github.com/wincentrtz/gobase/gobase/constants/noun"
	"strings"
)

func ConvertToPluralNoun(n string) string {
	for _, s := range noun.SuffixListWithEsReplacement {
		if strings.HasSuffix(n, s) {
			return n + "es"
		}
	}
	if strings.HasSuffix(n, noun.SuffixWithIesReplacement) {
		return n[:len(n)-1] + "ies"
	}

	if len(noun.IrregularNounWithCustomReplacement[n]) != 0 {
		return noun.IrregularNounWithCustomReplacement[n]
	}
	return n + "s"
}
