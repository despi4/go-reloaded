package autocorrection

import (
	"regexp"
	"strings"
)

var generalPattern = regexp.MustCompile(`(?i)(low|up|cap|hex|bin)`)

func Navigator(text string) string {
	flags := generalPattern.FindAllString(text, -1)

	// Functions to be applied based on flags
	flagFunc := map[string]func(string) string{
		"low": ToLower,
		"cap": FirstBig,
		"up":  ToUpper,
		"hex": HexToDec,
		"bin": BinToDec,
	}

	fixFunc := map[string]func(string) string{
		"low": AlphaNavigator,
		"cap": AlphaNavigator,
		"up":  AlphaNavigator,
		"hex": NumericNavigator,
		"bin": NumericNavigator,
	}

	// Apply initial text manipulations
	text = Insert(text)
	text = CutSpace(text)
	text = FixSpace(text)

	// Track the previous state of the text to prevent infinite loop
	var prevText string

	// Loop until no flags are found or the text stops changing
	for len(flags) != 0 {
		// If text is the same as the previous iteration, break out of the loop
		if text == prevText {
			break
		}
		// Save the current text to check if it changes in the next iteration
		prevText = text

		// Process the first flag
		str := flags[0]

		// Apply corresponding flag function if available
		function, ok := flagFunc[strings.ToLower(str)]
		if ok {
			text = function(text)
			text = FixSpace(text)
		}

		// Apply corresponding fix function if available
		fix, ok := fixFunc[strings.ToLower(str)]
		if ok {
			text = fix(text)
			text = FixSpace(text)
		}

		// Update flags based on modified text
		flags = generalPattern.FindAllString(text, -1)
	}

	// Final cleanup and return
	text = Punctuation(text)
	text = CutSpace(text)
	text = FixSpace(text)
	text = FixArticles(text)
	text = FixSpace(text)

	return text
}
