// gosplitargs
package gosplitargs

import (
	"regexp"
	"strings"
)

func SplitArgs(input, separator string) ([]string, error) {
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
		if element == "'" {
			if !doubleQuoteOpen {
				singleQuoteOpen = !singleQuoteOpen
				continue
			}
		} else if element == `"` {
			if !singleQuoteOpen {
				doubleQuoteOpen = !doubleQuoteOpen
				continue
			}
		}

		if !singleQuoteOpen && !doubleQuoteOpen {
			if matches {
				if len(tokenBuffer) > 0 {
					ret = append(ret, strings.Join(tokenBuffer, ""))
					tokenBuffer = make([]string, 0)
				}
			} else {
				tokenBuffer = append(tokenBuffer, element)
			}
		} else if singleQuoteOpen {
			tokenBuffer = append(tokenBuffer, element)
		} else if doubleQuoteOpen {
			tokenBuffer = append(tokenBuffer, element)
		}
	}
	if len(tokenBuffer) > 0 {
		ret = append(ret, strings.Join(tokenBuffer, ""))
	}
	return ret, nil
}
