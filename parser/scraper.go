package parser

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseHtmlContent get html from url and find all text content between each paragraph tags.
func ParseHtmlContent(url string) (value string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var content []string
	var totalLen int
	maxLen := 2000
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if len(text) > 0 {
			content = append(content, text)
			totalLen += len(text)
		}
	})

	fmt.Printf("HTML content length: %d\n", totalLen)

	var finLen int
	// split
	for i, text := range content {
		lText := len(text)
		if (lText + finLen) > maxLen {
			content = slices.Delete(content, i, i)
		} else {
			finLen += lText
		}
	}

	fmt.Printf("Final HTML content length: %d\n", finLen)

	return strings.Join(content, "\n\n"), nil
}
