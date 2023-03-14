package audit

import (

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


func scanHeadings(doc *html.Node, auditDoc *HtmlDocumentAudit) {
  var prevHeading *Heading 
  var crawler func(*html.Node)
  crawler = func(node *html.Node) {
    if node.Type == html.ElementNode && isHeadingElement(node.Data) {
      heading := newHeading(node.Data, getNodeText(node))
      auditDoc.SaveHeading(heading, prevHeading)
      prevHeading = heading 
    }

    for child := node.FirstChild; child != nil; child = child.NextSibling {
      crawler(child)
    }
  }

  crawler(doc)
}

