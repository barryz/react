// Code generated by reactGen; DO NOT EDIT.

package react

// H3Props defines the properties for the <h3> element
type H3Props struct {
	ID                      string
	Key                     string
	ClassName               string
	Role                    string
	OnChange                func(e *SyntheticEvent)
	OnClick                 func(e *SyntheticMouseEvent)
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
}

func (h *H3Props) assign(v *_H3Props) {

	if h.ID != "" {
		v.ID = h.ID
	}

	if h.Key != "" {
		v.Key = h.Key
	}

	v.ClassName = h.ClassName

	v.Role = h.Role

	v.OnChange = h.OnChange

	v.OnClick = h.OnClick

	v.DangerouslySetInnerHTML = h.DangerouslySetInnerHTML

}