package audit

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var headingTags = map[string]int {
  "h1": 1,
  "h2": 2,
  "h3": 3,
  "h4": 4,
  "h5": 5,
  "h6": 6,
}

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

type Heading struct {
  Tag string 
  Text string 
  Parent *Heading 
  Children []*Heading 
}

func (h *Heading) isParentHeading(heading *Heading) bool {
  item1 := headingTags[h.Tag]
  item2 := headingTags[heading.Tag]
  return item2 - item1 > 0
}

func newAuditHTMLDoc() *HtmlDocumentAudit {
  return &HtmlDocumentAudit{
    Headings: []*Heading{},
  }
}

func newHeading(tag, text string) *Heading {
  return &Heading{
    Tag: tag,
    Text: text,
    Children: []*Heading{},
  }
}

func isHeadingElement(tag string) bool {
  _, ok := headingTags[tag]
  return ok
}

func getHeadingText(node *html.Node) string {
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


func scanHeadings(doc *html.Node, auditDoc *HtmlDocumentAudit) {
  var prevHeading *Heading 
  var crawler func(*html.Node)
  crawler = func(node *html.Node) {
    if node.Type == html.ElementNode && isHeadingElement(node.Data) {
      heading := newHeading(node.Data, getHeadingText(node))
      auditDoc.saveHeading(heading, prevHeading)
      prevHeading = heading 
    }

    for child := node.FirstChild; child != nil; child = child.NextSibling {
      crawler(child)
    }
  }

  crawler(doc)
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
