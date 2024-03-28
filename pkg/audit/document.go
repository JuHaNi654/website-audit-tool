package audit

import "sync"

type HtmlDocumentAudit struct {
	Domain   string
	Headings []*Heading
	Links    []*LinkNode
	mu       sync.Mutex
}

func (d *HtmlDocumentAudit) SaveHeading(current, previous *Heading) {
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

func NewAuditHTMLDoc() *HtmlDocumentAudit {
	return &HtmlDocumentAudit{
		Headings: []*Heading{},
	}
}

func (d *HtmlDocumentAudit) UpdateLinkList(link *LinkNode) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Links = append(d.Links, link)
}
