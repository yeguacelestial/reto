package getlinks

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeguacelestial/reto/utils"
)

func TestParseLinksFromHtml(t *testing.T) {
	// Create struct for links
	var links []Link

	linkExample := Link{
		Href: "/other-page",
		Text: "A link to another page",
	}

	expectedLinks := append(links, linkExample)

	stringHtml := utils.ParseHtmlFromUrl("https://raw.githubusercontent.com/gophercises/link/master/ex1.html")
	stringReaderHtml := strings.NewReader(stringHtml)

	receivedLinks, err := ParseLinksFromHtmlReader(stringReaderHtml)

	fmt.Println(receivedLinks, err)

	assert.Equal(t, expectedLinks, receivedLinks, "Received unexpected content instead of links.")
}
