package types

type TemplateData map[string][]TemplateDataEntry

type TemplateDataEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NewMessage struct {
	Text         string       `json:"text,omitempty"`
	TemplateName string       `json:"template_name,omitempty"`
	TemplateData TemplateData `json:"template_data"`
}
