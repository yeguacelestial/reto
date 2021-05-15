package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests whether the HTML from an url is correctly parsed
func TestParseHtmlFromUrl(t *testing.T) {
	expectedHtml := `
	<html>
		<body>
			<h1>Hello!</h1>
			<a href="/other-page">A link to another page</a>
		</body>
	</html>
	`

	expectedHtml = RemoveEscapeSquences(expectedHtml)

	receivedHtml := ParseHtmlFromUrl("https://raw.githubusercontent.com/gophercises/link/master/ex1.html")
	receivedHtml = RemoveEscapeSquences(receivedHtml)

	assert.Equal(t, expectedHtml, receivedHtml, "Received unexpected HTML.")
}

// Removes escape sequences from string
func RemoveEscapeSquences(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, ">  <", "><", -1)

	return s
}
