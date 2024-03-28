package audit

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func initTestHtmlContent() *html.Node {
	b, _ := os.ReadFile("./test-html/index.html")
	doc, _ := html.Parse(strings.NewReader(string(b)))
	return doc
}

func generateTestResult() *HtmlDocumentAudit {
	doc := NewAuditHTMLDoc()
	item1 := newHeading("h1", "This is a h1 title With span")

	item2 := newHeading("h3", "Heading 1-3")
	item2.Parent = item1
	item1.Children = append(item1.Children, item2)

	item3 := newHeading("h4", "Heading 1-3-4")
	item3.Parent = item2
	item2.Children = append(item2.Children, item3)

	item4 := newHeading("h2", "Heading 1-2")
	item4.Parent = item1
	item1.Children = append(item1.Children, item4)

	item5 := newHeading("h5", "Heading 1-2-5")
	item5.Parent = item4
	item4.Children = append(item4.Children, item5)

	doc.Headings = append(doc.Headings, item1)

	return doc
}

func TestValidHeadingCrawl(t *testing.T) {
	doc := NewAuditHTMLDoc()
	data := initTestHtmlContent()
	scanHeadings(data, doc)

	testResult := generateTestResult()

	var checkMatchingHeadings func([]*Heading, []*Heading)
	checkMatchingHeadings = func(list1 []*Heading, list2 []*Heading) {
		for i, h := range list1 {
			r := list2[i]
			if h.Tag != r.Tag {
				t.Fatalf("Missmatching tags - Current %s - Should be %s", h.Tag, r.Tag)
			}

			if h.Text != r.Text {
				t.Fatalf("Missmatching text - Current %s - Should be %s", h.Tag, r.Tag)
			}

			checkMatchingHeadings(h.Children, r.Children)
		}
	}

	checkMatchingHeadings(testResult.Headings, doc.Headings)
}
