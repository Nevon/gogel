package gogel

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseElement(t *testing.T) {
	assert := assert.New(t)
	input := "<div id='foo' class='bar-baz'>Hello, 你叫什麼名字?</div>"
	tn := NewTextNode("Hello, 你叫什麼名字?")
	expected := NewElementNode("div", map[string]string{"id": "foo", "class": "bar-baz"}, []Node{tn})
	parser := NewParser(input)
	assert.Equal(expected, parser.parseElement())
}

func TestParseText(t *testing.T) {
	assert := assert.New(t)
	input := "Hello, 你叫什麼名字?</p>"
	expected := NewTextNode("Hello, 你叫什麼名字?")
	parser := NewParser(input)
	assert.Equal(expected, parser.parseText())
}

func TestParseAttributes(t *testing.T) {
	input := "id='foo' class='bar'>"
	expected := map[string]string{
		"id":    "foo",
		"class": "bar",
	}

	assert := assert.New(t)
	parser := NewParser(input)
	assert.Equal(expected, parser.parseAttributes())
}

func TestParseAttribute(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		in    string
		key   string
		value string
	}{
		{
			"id='foo'",
			"id",
			"foo",
		},
		{
			"class=\"testing\"",
			"class",
			"testing",
		},
		{
			"class='foo-bar_'",
			"class",
			"foo-bar_",
		},
	}

	for _, tt := range cases {
		parser := NewParser(tt.in)
		name, value := parser.parseAttribute()

		assert.Equal(tt.key, name)
		assert.Equal(tt.value, value)
	}
}

func TestStartsWith(t *testing.T) {
	assert := assert.New(t)
	input := "<div>"
	parser := NewParser(input)

	assert.True(parser.startsWith("<"), "%v starts with <", input)
	assert.True(parser.startsWith("<d"), "%v starts with <d", input)
	assert.True(parser.startsWith("<div>"), "%v starts with <div>", input)
	assert.False(parser.startsWith("g"), "%v does not start with g", input)
	assert.False(parser.startsWith(" "), "%v does not start with ' '", input)
}

func TestNextChar(t *testing.T) {
	assert := assert.New(t)
	input := "abc123字漢字"
	expected := strings.Split(input, "")
	parser := NewParser(input)

	for _, char := range expected {
		assert.Equal(char, parser.nextChar())
		parser.consumeChar()
	}
}

func TestEof(t *testing.T) {
	input := "a"
	parser := NewParser(input)
	assert.False(t, parser.eof())
	parser.consumeChar()
	assert.True(t, parser.eof())
}

func TestParseTagName(t *testing.T) {
	assert := assert.New(t)
	input := "<html id='2'>"
	parser := NewParser(input)
	parser.consumeChar()
	name := parser.parseTagName()

	assert.Equal("html", name)
}

func TestConsumeWhiteSpace(t *testing.T) {
	assert := assert.New(t)
	input := "      &%bla rgh"
	parser := NewParser(input)
	assert.Equal("      ", parser.consumeWhitespace())
}
