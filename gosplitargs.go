// gosplitargs
package gosplitargs

import (
	"regexp"
	"strings"
)

func SplitArgs(input, separator string, keepQuotes bool) ([]string, error) {
	if separator == "" {
		separator = "\\s"
	}
	singleQuoteOpen := false
	doubleQuoteOpen := false
	var tokenBuffer []string
	var ret []string

	arr := strings.Split(input, "")
	for _, element := range arr {
		matches, err := regexp.MatchString(separator, element)
		if err != nil {
			return nil, err
		}
		if element == "'" && !doubleQuoteOpen {
			if keepQuotes {
				tokenBuffer = append(tokenBuffer, element)
			}
			singleQuoteOpen = !singleQuoteOpen
			continue
		} else if element == `"` && !singleQuoteOpen {
			if keepQuotes {
				tokenBuffer = append(tokenBuffer, element)
			}
			doubleQuoteOpen = !doubleQuoteOpen
			continue
		}

		if !singleQuoteOpen && !doubleQuoteOpen && matches {
			if len(tokenBuffer) > 0 {
				ret = append(ret, strings.Join(tokenBuffer, ""))
				tokenBuffer = make([]string, 0)
			}
		} else {
			tokenBuffer = append(tokenBuffer, element)
		}
	}
	if len(tokenBuffer) > 0 {
		ret = append(ret, strings.Join(tokenBuffer, ""))
	}
	return ret, nil
}
