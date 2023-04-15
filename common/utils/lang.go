package utils

import (
	"errors"

	"golang.org/x/text/language"
)

var (
	errInvalidLanguage = errors.New("invalid language")
)

func ParseLanguageTag(lang string) (string, error) {
	tag, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		return "", err
	}
	if len(tag) == 0 {
		return "", errInvalidLanguage
	}
	t, _ := tag[0].Base()
	return t.String(), nil
}
