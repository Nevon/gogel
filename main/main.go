package main

import (
	"fmt"
	"github.com/nevon/gogel"
	"strings"
)

func main() {
	// body := gogel.NewElementNode("body", make(map[string]string), []gogel.Node{})
	// body.Children = append(body.Children, gogel.NewElementNode("div", map[string]string{"id": "test"}, []gogel.Node{}))

	// Walk(body, 0)

	doc := "<html><body><h1>Hello friends</h1><div id='myDiv'><p>This is a paragraph</p><p>Followed by another paragraph.</p></div></body></html>"

	tree := gogel.Parse(doc)
	Walk(tree[0], 0)
}

func Walk(n gogel.Node, level int) {
	out := []string{
		strings.Repeat("  ", level),
		n.ToString(),
	}
	fmt.Println(strings.Join(out, ""))

	for _, child := range n.GetChildren() {
		Walk(child, level+1)
	}
}
