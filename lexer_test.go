package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizeCaseSensitive(t *testing.T) {
	assertions := assert.New(t)

	tokenTypes := []TokenType{}
	tokenTypes = append(tokenTypes, TokenType{
		Type:          "keyword",
		Words:         []string{"if", "for"},
		CaseSensitive: true,
	})
	result := Tokenize("if", tokenTypes)
	assertions.Equal("if", result[0].Text)
	assertions.Equal(0, result[0].Position)
	assertions.Equal("keyword", result[0].Type)
}

func TestTokenizeCaseInsensitive(t *testing.T) {
	assertions := assert.New(t)

	tokenTypes := []TokenType{}
	tokenTypes = append(tokenTypes, TokenType{
		Type:          "keyword",
		Words:         []string{"if", "for"},
		CaseSensitive: false,
	})
	result := Tokenize("IF", tokenTypes)
	assertions.Equal("IF", result[0].Text)
	assertions.Equal(0, result[0].Position)
	assertions.Equal("keyword", result[0].Type)
}

func TestTokenizeWord(t *testing.T) {
	assertions := assert.New(t)

	tokenTypes := []TokenType{}
	tokenTypes = append(tokenTypes, TokenType{
		Type:          "keyword",
		Words:         []string{"if", "for"},
		CaseSensitive: true,
	})
	result := TokenizeWord("if", 0, tokenTypes)
	assertions.Equal("if", result.Text)
	assertions.Equal(0, result.Position)
	assertions.Equal("keyword", result.Type)
}

func TestLookupType(t *testing.T) {
	assertions := assert.New(t)

	tokenTypes := []TokenType{}
	tokenTypes = append(tokenTypes, TokenType{
		Type:          "keyword",
		Words:         []string{"if", "for"},
		CaseSensitive: false,
	})

	result := LookupType("for", tokenTypes)
	assertions.Equal("keyword", result.Type)
	assertions.Equal([]string{"if", "for"}, result.Words)
	assertions.Equal(false, result.CaseSensitive)
}