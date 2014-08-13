package gogel

import (
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {
	position int
	input    []string
}

func Parse(source string) []Node {
	p := NewParser(source)
	nodes := p.parseNodes()

	return nodes
}

func NewParser(source string) *Parser {
	return &Parser{
		position: 0,
		input:    strings.Split(source, ""),
	}
}

// Traverses a new node down the tree
func (p *Parser) parseNodes() []Node {
	nodes := []Node{}

	for {
		p.consumeWhitespace()

		if p.eof() || p.startsWith("</") {
			break
		}

		nodes = append(nodes, p.parseNode())
	}

	return nodes
}

// Parses a single text or element node
func (p *Parser) parseNode() Node {
	if p.nextChar() == "<" {
		return p.parseElement()
	} else {
		return p.parseText()
	}
}

func (p *Parser) parseText() *TextNode {
	return NewTextNode(p.consumeWhile(func(char string) bool {
		return char != ">"
	}))
}

func (p *Parser) parseElement() *ElementNode {
	/* This is a clusterfuck of panics. Please give me assertions! */
	start := p.consumeChar()
	if start != "<" {
		panic(fmt.Sprintf("Malformed input. %v was not an opening <", start))
	}

	tagName := p.parseTagName()
	attributes := p.parseAttributes()

	end := p.consumeChar()
	if end != ">" {
		panic(fmt.Sprintf("Malformed input. %v was not a closing >", end))
	}

	children := p.parseNodes()

	startEndTag := p.consumeChar()
	endTagSlash := p.consumeChar()

	if startEndTag != "<" {
		panic(fmt.Sprintf("Malformed input. %v was not an opening <", startEndTag))
	}

	if endTagSlash != ">" {
		panic(fmt.Sprintf("Malformed input. %v was not a closing >", endTagSlash))
	}

	endTagName := p.parseTagName()

	if endTagName != tagName {
		panic(fmt.Sprintf("Malformed input. Closing %v tag does not match opening %v", endTagName, tagName))
	}

	endTagEnd := p.consumeChar()
	if endTagEnd != ">" {
		panic(fmt.Sprintf("Malformed input. %v was not a closing >", endTagEnd))
	}

	return NewElementNode(tagName, attributes, children)
}

func (p *Parser) parseAttribute() (string, string) {
	name := p.parseTagName()
	delimiter := p.consumeChar()

	if delimiter != "=" {
		panic(fmt.Sprintf("Malformed input. %v was not a = in attribute-value pair at index %v in %v", delimiter, p.position, p.input))
	}

	value := p.parseAttributeValue()
	return name, value
}

func (p *Parser) parseAttributeValue() string {
	openQuote := p.consumeChar()

	if openQuote != "\"" && openQuote != "'" {
		panic(fmt.Sprintf("Malformed input. %v was not an opening quote", openQuote))
	}

	value := p.consumeWhile(func(char string) bool {
		return char != openQuote
	})

	closeQuote := p.consumeChar()
	if closeQuote != openQuote {
		panic(fmt.Sprintf("Malformed input. Closing %v was not the same as opening %v", closeQuote, openQuote))
	}

	return value
}

func (p *Parser) parseAttributes() map[string]string {
	attributes := map[string]string{}

	for {
		p.consumeWhitespace()
		if p.nextChar() == ">" {
			break
		}
		name, value := p.parseAttribute()
		attributes[name] = value
	}

	return attributes
}

func (p *Parser) nextChar() string {
	return p.input[p.position]
}

func (p *Parser) startsWith(needle string) bool {
	return p.input[0] == needle
}

func (p *Parser) eof() bool {
	return p.position >= len(p.input)
}

func (p *Parser) consumeChar() string {
	char := p.input[p.position]
	p.position++
	return char
}

func (p *Parser) consumeWhile(test func(char string) bool) string {
	result := []string{}

	for {
		if p.eof() || !test(p.nextChar()) {
			break
		}
		result = append(result, p.consumeChar())
	}

	return strings.Join(result, "")
}

func (p *Parser) consumeWhitespace() {
	p.consumeWhile(func(char string) bool {
		return char == " " || char == "\n" || char == "\t" || char == "\r"
	})
}

func (p *Parser) parseTagName() string {
	reg := regexp.MustCompile("[A-Za-z0-9]")

	return p.consumeWhile(func(char string) bool {
		bytes := []byte(char)
		return reg.Match(bytes)
	})
}
