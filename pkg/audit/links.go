package audit

import (
	"fmt"

	"golang.org/x/net/html"
)

type Link struct {
  StatusCode int8
  Label string
  Type string
  Link string 
}

func scanLinks(doc *html.Node, auditDoc *HtmlDocumentAudit) {
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
      fmt.Println()
    }
	}

  crawler(doc)
}
