package audit

import (
	"strings"

	vs "github.com/JuHaNi654/website-audit-tool/pkg/constant"
	"golang.org/x/net/html"
)

func getNodeText(node *html.Node) string {
	text := ""
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode {
			text += node.Data
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(node)
	return strings.TrimSpace(text)
}

func RunAudit(url string, action vs.ScanAction) (*HtmlDocumentAudit, error) {
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

	if action == vs.ScanLinks {
		scanLinks(doc, auditDoc)
	} else if action == vs.ScanHeading {
		scanHeadings(doc, auditDoc)
	}

	return auditDoc, nil
}
