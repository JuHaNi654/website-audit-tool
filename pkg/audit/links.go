package audit

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type LinkNode struct {
	StatusCode uint16
	Label      string
	Type       string
	Link       string
	Message    string
}

func newLinkNode(label, link string) *LinkNode {
	return &LinkNode{
		Label: label,
		Link:  link,
	}
}

func checkLinkType(attr []html.Attribute) string {
	for _, item := range attr {
		if item.Key == "href" {
			return item.Val
		}
	}

	return ""
}

// TODO check if there are multiple in same id targets
func findAnchorTarget(target string, node *html.Node) uint16 {
	status := 404
	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, i := range n.Attr {
				if i.Key == "id" && i.Val == target {
					status = 200
					return
				}
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(node)
	return uint16(status)
}

func checkValidLinks(l *LinkNode, domain string, node *html.Node) {
	if l.Link == "" {
		l.Message = "Inactive link tag"
		return
	}

	if strings.HasPrefix(l.Link, "tel:") || strings.HasPrefix(l.Link, "mailto:") {
		l.Message = "Contact link"
		return
	}

	if strings.HasPrefix(l.Link, "#") {
		if len(l.Link) == 1 {
			l.Message = "Inactive anchor link"
		} else {
			// TODO check if anchor target is found from document
			l.Message = "Anchor link"
			mc := regexp.MustCompile(`^#`)
			target := mc.ReplaceAllString(l.Link, "")
			l.StatusCode = findAnchorTarget(target, node)

			if l.StatusCode == 404 {
				l.Message = "Anchor link target not found"
			}
		}
		return
	}

	// Combine relative path links with domain
	if strings.HasPrefix(l.Link, "/") || strings.HasPrefix(l.Link, "./") {
		mc := regexp.MustCompile(`^.`)
		path := mc.ReplaceAllString(l.Link, "")
		l.Link = fmt.Sprintf("%s/%s", domain, path)
	}

	// TODO should application check different domain links
	status, err := checkActiveUrl(l.Link)
	if err != nil {
		l.Message = err.Error()
	}
	l.StatusCode = status
}

func scanLinks(doc *html.Node, auditDoc *HtmlDocumentAudit) {
	var crawler func(*html.Node, *sync.WaitGroup)
	var wg sync.WaitGroup
	crawler = func(node *html.Node, wg *sync.WaitGroup) {
		defer wg.Done()
		if node.Type == html.ElementNode && node.Data == "a" {
			newLink := newLinkNode(getNodeText(node), checkLinkType(node.Attr))
			checkValidLinks(newLink, auditDoc.Domain, doc)
			auditDoc.UpdateLinkList(newLink)
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			wg.Add(1)
			go crawler(child, wg)
		}

	}
	wg.Add(1)
	go crawler(doc, &wg)
	wg.Wait()
}
