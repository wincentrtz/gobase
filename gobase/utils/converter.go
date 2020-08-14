package utils

import "strings"

func ConvertToPluralNoun(noun string) string {
	suffixList := []string{
		"s", "x", "z", "ch", "sh",
	}
	for _, s := range suffixList {
		if strings.HasSuffix(noun, s) {
			return noun + "es"
		}
	}
	if strings.HasSuffix(noun, "y") {
		return noun[:len(noun)-1] + "ies"
	}

	irregularNoun := map[string]string{
		"woman":      "women",
		"man":        "men",
		"child":      "children",
		"tooth":      "teeth",
		"foot":       "feet",
		"person":     "people",
		"leaf":       "leaves",
		"mouse":      "mice",
		"goose":      "geese",
		"half":       "halves",
		"knife":      "knives",
		"wife":       "wives",
		"life":       "lives",
		"elf":        "elves",
		"loaf":       "loaves",
		"potato":     "potatoes",
		"tomato":     "tomatoes",
		"cactus":     "cacti",
		"focus":      "foci",
		"fungus":     "fungi",
		"nucleus":    "nuclei",
		"syllabus":   "syllabuses",
		"analysis":   "analyses",
		"diagnosis":  "diagnoses",
		"oasis":      "oases",
		"thesis":     "theses",
		"crisis":     "crises",
		"phenomenon": "phenomena",
		"criterion":  "criteria",
		"datum":      "data",
	}
	if len(irregularNoun[noun]) != 0 {
		return irregularNoun[noun]
	}
	return noun + "s"
}
