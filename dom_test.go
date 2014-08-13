package gogel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElementNodeToString(t *testing.T) {
	assert := assert.New(t)
	tagName := "div"
	attributes := map[string]string{}
	children := []Node{}

	element := NewElementNode(tagName, attributes, children)

	assert.Equal(tagName, element.ToString())
}

func TestElementNodeGetChildren(t *testing.T) {
	assert := assert.New(t)
	tagName := "p"
	attributes := map[string]string{}
	children := []Node{
		NewTextNode("Hello"),
	}

	element := NewElementNode(tagName, attributes, children)

	assert.Equal(children, element.GetChildren())
}

func TestTextNodeToString(t *testing.T) {
	assert := assert.New(t)
	text := "שלום, חבר שלי!"

	element := NewTextNode(text)

	assert.Equal(text, element.ToString())
}
