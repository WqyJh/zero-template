package utils_test

import (
	"testing"
	"zero-template/common/utils"

	"github.com/stretchr/testify/assert"
)

func TestParseLang(t *testing.T) {
	var testData = []struct {
		lang string
		want string
	}{
		{"zh-Hans-CN;q=1.0, en-CN;q=0.9", "zh"},
		{"zh-Hans-CN;q=1.0", "zh"},
		{"zh-Hans-CN", "zh"},
		{"zh-Hans", "zh"},
		{"zh", "zh"},
		{"en", "en"},
		{"en-US", "en"},
		{"en-US;q=1.0, zh-Hans-CN;q=0.9", "en"},
		{"en-US;q=1.0, zh-Hans-CN;q=0.9, zh-Hans;q=0.8", "en"},
		{"en-US;q=1.0, zh-Hans-CN;q=0.9, zh-Hans;q=0.8, zh;q=0.7", "en"},
		{"en-US;q=1.0, zh-Hans-CN;q=0.9, zh-Hans;q=0.8, zh;q=0.7, ja;q=0.6", "en"},
	}
	for _, tt := range testData {
		lang, err := utils.ParseLanguageTag(tt.lang)
		assert.NoError(t, err)
		assert.Equal(t, tt.want, lang)
	}
}
