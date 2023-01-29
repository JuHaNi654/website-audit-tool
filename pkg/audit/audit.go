package audit

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)
type HtmlDocumentAudit struct {
  Headings []*Heading 
}

func (d *HtmlDocumentAudit) saveHeading(current, previous *Heading) {
  parent := previous 
  for parent != nil {
    if parent.isParentHeading(current) {
      current.Parent = parent 
      parent.Children = append(parent.Children, current)
      return 
    }

    parent = parent.Parent  
  }

  if parent == nil {
    d.Headings = append(d.Headings, current)
  }
}

func newAuditHTMLDoc() *HtmlDocumentAudit {
  return &HtmlDocumentAudit{
    Headings: []*Heading{},
  }
}


func getNodeText(node *html.Node) string {
  text := ""
  var crawler func(*html.Node)
  crawler = func(node *html.Node) {
    if node.Type == html.TextNode {
      text += fmt.Sprintf(" %s", node.Data)
    }

    for child := node.FirstChild; child != nil; child = child.NextSibling {
      crawler(child)
    }
  }
  crawler(node)
  return strings.TrimSpace(text)
}

func RunAudit(url string) (*HtmlDocumentAudit, error) {
  err := ValidateUrl(url)
  if err != nil {
    return nil, err
  }


  auditDoc := newAuditHTMLDoc()
  body, err := FetchPageDocument(url)
  if err != nil {
    return nil, err
  }

  doc, _ := html.Parse(strings.NewReader(string(body)))
  scanHeadings(doc, auditDoc)

  return auditDoc, nil
}
