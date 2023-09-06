package main

import "html/template"

var (
	tmpl *template.Template
)

func ParseTemplate(name string, content string) {
	if tmpl == nil {
		tmpl = template.New(name)
	}
	var t *template.Template
	if tmpl.Name() == name {
		t = tmpl
	} else {
		t = tmpl.New(name)
	}
	template.Must(t.New(name).Delims("[[", "]]").Parse(content))
}
