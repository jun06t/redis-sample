package autocomplete

import (
	"strings"
)

const validChars = "`abcdefghijklmnopqrstqvwxyz{"

func FindPrefixRange(prefix string) (start string, end string) {
	pos := strings.Index(validChars, prefix[len(prefix)-1:])
	suffix := validChars[pos-1 : pos]

	return prefix[:len(prefix)-1] + suffix + "{", prefix + "{"
}
