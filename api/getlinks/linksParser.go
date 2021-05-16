package getlinks

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Represents a link in an HTML document
type Link struct {
	Text string
	Href string
}

// Parse will take an HTML document and will return a slice of
// links parsed from it.
func ParseLinksFromHtmlReader(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)

	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	// Iterate through all the attributes of the node
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	// If the node is a text, return its content
	if n.Type == html.TextNode {
		return n.Data
	}

	// If the node is not an html element, return an empty string
	if n.Type != html.ElementNode {
		return ""
	}

	// Iterate the child elements from the node, and use recursion for
	// extracting the text
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}

	// Remove unnecessary spaces, and join words with a single space
	// between them
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	// Verifiy that the node is a valid anchor tag
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	// If the node is not anchor tag, still verifiy the child elements
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
