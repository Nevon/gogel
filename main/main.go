package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/nevon/gogel"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "gogel"
	app.Usage = "Parse an HTML document"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "html",
			Usage: "path to HTML file to parse",
		},
	}

	app.Action = func(c *cli.Context) {
		html, err := ioutil.ReadFile(c.String("html"))

		if err == nil {
			tree := gogel.Parse(string(html))
			Walk(tree[0], 0)
		}
	}

	app.Run(os.Args)
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
