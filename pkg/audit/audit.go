package audit

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// FIXME: trim extra spacing
func getNodeText(node *html.Node) string {
	text := ""
	space := regexp.MustCompile(`\s+`)
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode && node.Data != "" {
			text += strings.TrimSpace(node.Data)
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(node)
	return space.ReplaceAllString(text, " ")
}

func RunAudit(url string, action ScanAction) (*HtmlDocumentAudit, error) {
	err := ValidateUrl(url)
	if err != nil {
		return nil, err
	}

	auditDoc := NewAuditHTMLDoc()
	auditDoc.Domain = getDomain(url)

	// TODO handle errors
	body, err := fetchPageDocument(url)
	if err != nil {
		return nil, err
	}

	doc, _ := html.Parse(strings.NewReader(string(body)))

	if action == ScanLinks {
		scanLinks(doc, auditDoc)
	} else if action == ScanHeading {
		scanHeadings(doc, auditDoc)
	}

	return auditDoc, nil
}
