// gosplitargs
package gosplitargs

import (
	"strings"
)

func SplitArgs(input string, separator string, keepQuotes bool) ([]string, error) {
	return splitArgs(input, separator, keepQuotes, "")
}

func SplitSQL(input string, separator string, keepQuotes bool) ([]string, error) {
	return splitArgs(input, separator, keepQuotes, "--")
}

func splitArgs(input string, separator string, keepQuotes bool, commentSign string) ([]string, error) {
	singleQuoteOpen := false
	doubleQuoteOpen := false
	commentSignOpen := false
	commentSignMatchingIndex := 0
	separatorMatchingIndex := 0

	var tokenBuffer string
	var ret []string

	inputCharArray := strings.Split(input, "")
	for _, inputChar := range inputCharArray {
		// special cases
		if inputChar == "\n" {
			commentSignMatchingIndex = 0
			commentSignOpen = false
			tokenBuffer += inputChar
			continue
		}

		if commentSignOpen {
			tokenBuffer += inputChar
			continue
		}

		if inputChar == "'" && !doubleQuoteOpen && !commentSignOpen {
			if keepQuotes {
				tokenBuffer += inputChar
			}
			singleQuoteOpen = !singleQuoteOpen
			continue
		} else if inputChar == `"` && !singleQuoteOpen && !commentSignOpen {
			if keepQuotes {
				tokenBuffer += inputChar
			}
			doubleQuoteOpen = !doubleQuoteOpen
			continue
		}

		// normal cases
		if !singleQuoteOpen && !doubleQuoteOpen && !commentSignOpen {
			if commentSign != "" {
				if inputChar == string(commentSign[commentSignMatchingIndex]) {
					// it could be a comment
					commentSignMatchingIndex++
					if commentSignMatchingIndex == len(commentSign) {
						// it's a comment
						commentSignOpen = true
						commentSignMatchingIndex = 0
					}
					tokenBuffer += inputChar
					continue
				} else {
					// it's not a comment
					commentSignMatchingIndex = 0
				}
			}

			if separator == "" {
				if strings.TrimSpace(inputChar) == "" {
					// it's a separator
					if len(tokenBuffer) > 0 {
						ret = append(ret, tokenBuffer)
					}
					tokenBuffer = ""
				} else {
					// it's not a separator
					tokenBuffer += inputChar
				}
				continue
			} else if inputChar == string(separator[separatorMatchingIndex]) {
				// it could be a separator
				separatorMatchingIndex++
				if separatorMatchingIndex == len(separator) {
					// it's a separator
					ret = append(ret, tokenBuffer[:len(tokenBuffer)-len(separator)+1])
					tokenBuffer = ""
					separatorMatchingIndex = 0
				} else {
					// it's not a separator
					tokenBuffer += inputChar
				}
				continue
			} else {
				// it's not a separator
				separatorMatchingIndex = 0
			}
		}
		tokenBuffer += inputChar
	}

	if len(tokenBuffer) > 0 || strings.TrimSpace(separator) != "" {
		ret = append(ret, tokenBuffer)
	}
	return ret, nil
}

// func CountSeparators(input string, separator string) (int, error) {
// 	if separator == "" {
// 		separator = "\\s+"
// 	}
// 	singleQuoteOpen := false
// 	doubleQuoteOpen := false
// 	ret := 0

// 	arr := strings.Split(input, "")
// 	for _, element := range arr {
// 		matches, err := regexp.MatchString(separator, element)
// 		if err != nil {
// 			return -1, err
// 		}
// 		if element == "'" && !doubleQuoteOpen {
// 			singleQuoteOpen = !singleQuoteOpen
// 			continue
// 		} else if element == `"` && !singleQuoteOpen {
// 			doubleQuoteOpen = !doubleQuoteOpen
// 			continue
// 		}

// 		if !singleQuoteOpen && !doubleQuoteOpen && matches {
// 			ret++
// 		}
// 	}
// 	return ret, nil
// }
