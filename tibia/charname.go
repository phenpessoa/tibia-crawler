package tibia

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsCharNameValid checks if the provided name is a valid character name
// according to the rules imposed by tibia.com.
// The validation can be performed based on either the new rules or the legacy
// rules, depending on the value of the newRules boolean.
//
// If newRules is set to true, the function validates the name based on the new
// rules. These rules represent the current character name restrictions imposed
// by tibia.com. It ensures that the name satisfies the current requirements for
// creating or renaming a character.
//
// If newRules is set to false, the function validates the name based on the
// legacy rules. Although legacy names still exist in the system, they are not
// used for creating or renaming characters anymore. The function checks if the
// name would have been valid under the old rules.
//
// For example, Torbjörn used to be a valid name, and this character still
// exists, but it nos a valid name by the new rules anymore.
func IsCharNameValid(name string, newRules bool) bool {
	if name == "" {
		return false
	}

	trimmed := strings.TrimSpace(name)
	if trimmed == "" || trimmed != name {
		return false
	}

	runeCount := utf8.RuneCountInString(name)
	if runeCount > MaxRunesAllowedInCharName ||
		runeCount < MinRunesAllowedInCharName {
		return false
	}

	fields := strings.Fields(name)
	for _, field := range fields {
		runeCount := utf8.RuneCountInString(field)
		if runeCount > MaxRunesAllowedInCharWord ||
			runeCount < MinRunesAllowedInCharWord {
			return false
		}

		if newRules && runeCount > MaxRunesAllowedInCharWordNew {
			return false
		}
	}

	var isLastRuneSpace bool
	for _, r := range name {
		if !isRuneValidCharRune(r, newRules) {
			return false
		}

		if r == ' ' && isLastRuneSpace {
			return false
		}

		if r == ' ' {
			isLastRuneSpace = true
		} else {
			isLastRuneSpace = false
		}
	}

	return true
}

func isRuneValidCharRune(r rune, newRules bool) bool {
	if newRules {
		return isRuneValidCharRuneNew(r)
	}
	return isRuneValidCharRuneOld(r)
}

func isRuneValidCharRuneNew(r rune) bool {
	if (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') {
		return true
	}

	if r == ' ' {
		return true
	}

	return false
}

func isRuneValidCharRuneOld(r rune) bool {
	switch {
	case unicode.IsLetter(r):
		return true
	case r == ' ':
		return true
	case r == '\'':
		return true
	case r == '-':
		return true
	default:
		return false
	}
}
