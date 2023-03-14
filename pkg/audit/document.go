package audit

type HtmlDocumentAudit struct {
	Domain   string
	Headings []*Heading
	Links		[]*LinkNode
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
