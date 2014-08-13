package gogel

type Node interface {
	ToString() string
	GetChildren() []Node
}

type ElementNode struct {
	Children   []Node
	TagName    string
	Attributes map[string]string
}

func (n *ElementNode) ToString() string {
	return n.TagName
}

func (n *ElementNode) GetChildren() []Node {
	return n.Children
}

func NewElementNode(tagName string, attributes map[string]string, children []Node) *ElementNode {
	return &ElementNode{
		TagName:    tagName,
		Attributes: attributes,
		Children:   children,
	}
}

type TextNode struct {
	Text string
}

func (n *TextNode) ToString() string {
	return n.Text
}

func (n *TextNode) GetChildren() []Node {
	return []Node{}
}

func NewTextNode(text string) *TextNode {
	return &TextNode{
		Text: text,
	}
}
