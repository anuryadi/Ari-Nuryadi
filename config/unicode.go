package config

import (
	"strconv"
	"strings"
)

func UnescapeUnicodeCharactersInJSON(jsonRaw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
